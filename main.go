package main

import (
	"log"
	"github.com/gdamore/tcell/v2"
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

	if err !=nil {
		log.Fatalf("%+v",err)
	}

	initError:= screen.Init()
	width,height := screen.Size()

	if initError != nil {
		log.Fatalf("%+v",initError)
	}

	//Set default text style
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
	screen.SetStyle(defaultStyle)


	defer quit(screen)


	ball :=Ball{
		X:width/2,
		Y:height/2,
		SpeedX: 1,
		SpeedY: 1,
		Active:true,
		Screen: screen,
	}

	paddle:=Paddle{
		X: width/2,
		Y: height - 1,
		SpeedX: 10,
		Width: 20,
		Height: 1,
		Active:true,
	}

	var gameObjects []GameObject

	gameObjects = append(gameObjects,&ball,&paddle)


	blockWidth := 8
	blockHeight :=1
	blockOffset := 1
	initYPos:=10
	totalColumns:= (width / (blockWidth + blockOffset)) + blockWidth

	for i := range totalColumns{

		xPos :=0
		if (i > 0){
			xPos = ( (blockWidth + blockOffset) * i)
		}
		for j := range 7 {
			
			yPos:=initYPos
			if ( j > 0){
				yPos = initYPos+( (blockHeight + blockOffset) * j)
			}

			block :=Block{
				Width: blockWidth,
				Height: blockHeight,
				X:xPos,
				Y:yPos,
				Style: tcell.StyleDefault.Background(tcell.ColorOrange.TrueColor()).Foreground(tcell.ColorOrange.TrueColor()),
				Active:true,
			}
			gameObjects = append(gameObjects,&block)
		}
	}

	player := Player{
		Life: 3,
		Score: 0,
	}
	engine :=Engine{
		Screen:      screen,
		Style:       defaultStyle,
		GameObjects: gameObjects,
		Player: &player,
	}

	engine.InitEventListeners()
	go engine.Run()

	getUserInput(screen,&paddle)
 
}