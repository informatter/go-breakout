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
func update(engine *Engine){

	for _, gameObject := range engine.GameObjects{
		gameObject.Update()
		engine.Screen.SetContent(gameObject.GetX(), gameObject.GetY(), gameObject.Display(), nil, engine.Style)
	}
}

//Go routine which runs the game
func (engine *Engine) Run(){

	screen := engine.Screen

	for {
		screen.Clear()

		update(engine)

		time.Sleep(60 * time.Millisecond)

		screen.Show()
	}
}