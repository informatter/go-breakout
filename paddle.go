package main

import (
	"github.com/gdamore/tcell/v2"
	"strings"
)

type Paddle struct {
	Width  int
	Height int
	X      int
	Y      int
	SpeedX int
	Active bool
}

func (paddle Paddle) GetX() int{
	return paddle.X
}

func (paddle Paddle) GetActiveState() bool{
	return paddle.Active
}

func (paddle Paddle) GetY() int{
	return paddle.Y
}

func (paddle Paddle) GetHeight() int{
	return paddle.Height
}

func (paddle Paddle) GetWidth() int{
	return paddle.Width
}

func (paddle Paddle) Display(engine *Engine) {
	shape := strings.Repeat(" ", paddle.Width)
	paddleStyle := tcell.StyleDefault.Background(tcell.ColorPurple).Foreground(tcell.ColorPurple)

	renderGameObject(
		engine.Screen,
		paddle.GetX(),
		paddle.GetY(),
		paddle.GetX() + paddle.Width,
		paddle.GetY() + paddle.Height,
		paddleStyle,
		shape,
	)
}



func (paddle *Paddle) Update(engine Engine){

    select {
    case msg := <-PaddleMoveMessages:

        switch msg {
        case "paddle-move-left":
			paddle.X += paddle.SpeedX *-1

        case "paddle-move-right":
			paddle.X += paddle.SpeedX
        default:
        }
    default:

    }
}


func (paddle *Paddle) CheckEdges(screenWidth int, screenHeight int){
	
	if (paddle.X <= 0){
		paddle.X = 0

		//wrap around
		//paddle.X  = (screenWidth - paddle.Width ) - 1
	}
	
	if(paddle.X + paddle.Width >= screenWidth ){

		//wrap around
		//paddle.X  = (paddle.X + paddle.Width) - screenWidth 

		paddle.X = (screenWidth - paddle.Width)
	}

}

func (paddle *Paddle) CheckCollision(gameObject GameObject){

}

func (paddle *Paddle) Move(movementType string){
	
	if (movementType == "a"){
		// move left
		paddle.X += paddle.SpeedX *-1
	}else{
		// move right 
		paddle.X += paddle.SpeedX
	}
}

