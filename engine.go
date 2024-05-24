package main

import (
	"time"
	"github.com/gdamore/tcell/v2"
	"fmt"
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
}


//Go routine which runs the game
func (engine *Engine) Run(){

	screen := engine.Screen
	
	width,height := screen.Size()

	for {
		screen.Clear()
		renderGameObject(engine.Screen, 1, 1, 10, 1, engine.Style, fmt.Sprintf("Score: %d", engine.Player.Score))
		renderGameObject(engine.Screen, 1, 4, 10, 4, engine.Style, fmt.Sprintf("Life: %d", engine.Player.Life))

		update(engine,width,height)

		time.Sleep(20 * time.Millisecond)

		screen.Show()
	}
}