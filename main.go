package main

import (
	"log"
	//"os"
	"github.com/gdamore/tcell/v2"
	//"fmt"
)


/*
Polls for user input events and blocks the main thread while waiting 
*/
func getUserInput(screen tcell.Screen, paddle *Paddle ){ //eventChan chan<- Event
	for {
		event := screen.PollEvent() // waits for events for arrive and blocks the main thread

		switch event := event.(type){

		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				return
			}else if event.Rune() == 'a' {
				//game.Player1.Paddle.MoveUp()
				// TODO emmit event 
				// eventChan <- Event{
				// 	Value: "w",
				// }
				paddle.Move("a")
				//paddle.X += paddle.SpeedX 
			} else if event.Rune() == 'd' {
				// eventChan <- Event{
				// 	Value: "d",
				// }
				paddle.Move("d")
			}
		}
	}
}

func quit(screen tcell.Screen){
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		screen.Fini() //closes screen and releases resources
		if maybePanic != nil {
			panic(maybePanic)
		}
}

func main(){
	screen, err := tcell.NewScreen()

	//eventChan := make(chan Event)


	if err !=nil {
		log.Fatalf("%+v",err)
	}

	initError:= screen.Init()
	width,height := screen.Size()

	if initError != nil {
		log.Fatalf("%+v",initError)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
	screen.SetStyle(defStyle)


	defer quit(screen)


	ball :=Ball{
		X:width/2,
		Y:height/2,
		SpeedX: 1,
		SpeedY: 1,
	}

	paddle:=Paddle{
		X: width/2,//width/2 ,
		Y: height - 1,
		SpeedX: 2,
		Width: 20,
		Height: 1,
	}
	// ballB :=Ball{
	// 	X:0,
	// 	Y:0,
	// 	SpeedX: 0,
	// 	SpeedY: 1,
	// }
	engine :=Engine{
		Screen:      screen,
		Style:       defStyle,
		GameObjects: [] GameObject{&ball, &paddle}, //, &ballB},
	}
	go engine.Run()

	getUserInput(screen,&paddle)

	//paddle.HandleEvent(eventChan)
 
}