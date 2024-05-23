package main

import (
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
}



// Updates the state of all game objects.
func update(engine *Engine, screenWidth int, screenHeight int){

	for _, gameObject := range engine.GameObjects{

		gameObject.Update()
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

		update(engine,width,height)

		time.Sleep(20 * time.Millisecond)

		screen.Show()
	}
}