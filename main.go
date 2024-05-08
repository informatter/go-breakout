package main

import (
	"log"
	"os"
	"github.com/gdamore/tcell/v2"
)


/*
Polls for user input events and blocks the main thread while waiting 
*/
func getUserInput(screen tcell.Screen){
	for {
		event := screen.PollEvent() // waits for events for arrive and blocks the main thread

		switch event := event.(type){

		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini() // closes screen and releases resources
				os.Exit(0) // exits the program without error
			}
		}
	}
}

func main(){
	screen, err := tcell.NewScreen()

	if err !=nil {
		log.Fatalf("%+v",err)
	}

	initError:= screen.Init()

	if initError != nil {
		log.Fatalf("%+v",initError)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
	screen.SetStyle(defStyle)


	ballA :=Ball{
		X:1,
		Y:4,
		SpeedX: 1,
		SpeedY: 0,
	}
	ballB :=Ball{
		X:0,
		Y:0,
		SpeedX: 0,
		SpeedY: 1,
	}
	engine :=Engine{
		Screen:      screen,
		Style:       defStyle,
		GameObjects: [] GameObject{&ballA, &ballB},
	}
	go engine.Run()

	getUserInput(screen)
 
}