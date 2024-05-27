package main

import (
	"fmt"
	"time"
	"github.com/gdamore/tcell/v2"
)

/*
	Represents the game engine which is repsonsible to update all the game objects
	and render them to the screen.
*/
type Engine struct {
	Screen tcell.Screen
	Style tcell.Style
	GameObjects [] GameObject
	Score int
	Player *Player
	IsGameOver bool
}


// Updates the state of all game objects.
func update(engine *Engine, screenWidth int, screenHeight int){

	// handle collisions
	for i := range engine.GameObjects{

		current := engine.GameObjects[i]
		for j := range engine.GameObjects{

			if (i==j){continue}

			next := engine.GameObjects[j]

			current.CheckCollision(next)
		
		}
	}

	for _, gameObject := range engine.GameObjects{

		gameObject.Update(*engine)
		gameObject.CheckEdges(screenWidth,screenHeight)
		gameObject.Display(engine)

	}
	engine.receiveMessage()
}


func (engine Engine) receiveMessage() {


    select {
    case msg := <-GameStateMessages:
        switch msg {
        case "game-over":
			heightOffset := 4
			width,height := engine.Screen.Size()
			renderGameObject(engine.Screen, width/2, height/2 + heightOffset, 500, height/2 + heightOffset, engine.Style, "GAME OVER!")
			renderGameObject(engine.Screen, width/2 -3, (height/2) + heightOffset*2, 1000, (height/2) + heightOffset*2, engine.Style, "PRESS ESC TO EXIT")
        case "ball-droped":
			engine.Player.Life-=1
			renderGameObject(engine.Screen, 1, 4, 10, 4, engine.Style, fmt.Sprintf("Life: %d", engine.Player.Life))
        default:
        }
    default:

    }
}


//Go routine which runs the game
func (engine *Engine) Run(){

	screen := engine.Screen
	
	width,height := screen.Size()

	for  {
		
		screen.Clear()
		renderGameObject(engine.Screen, 1, 1, 10, 1, engine.Style, fmt.Sprintf("Score: %d", engine.Player.Score))
		renderGameObject(engine.Screen, 1, 4, 10, 4, engine.Style, fmt.Sprintf("Life: %d", engine.Player.Life))
		update(engine,width,height)
		
		time.Sleep(20 * time.Millisecond)
		screen.Show()
	}
	
}