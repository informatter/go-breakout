
package main


// Ball entity which implements the GameObject interface
type Ball struct{
	X int
	Y int
	SpeedX int
	SpeedY int
}

func (ball Ball) GetX() int{
	return ball.X
}

func (ball Ball) GetY() int{
	return ball.Y
}

func (ball Ball) GetHeight() int{
	return ball.GetY()
}

func (ball Ball) GetWidth() int{
	return ball.GetX()
}

func (ball Ball) Display(engine *Engine) {
	shape := "\u26AA"
	renderGameObject(
		engine.Screen,
		ball.GetX(),
		ball.GetY(),
		ball.GetX(),
		ball.GetY(),
		engine.Style,
		shape,
	)
}

func (ball *Ball) Update(){

	ball.X+=ball.SpeedX
	ball.Y+=ball.SpeedY

}

func (ball *Ball)  CheckCollision(g GameObject){



	hasCollision := ball.GetX() >= g.GetX() && ball.GetX() <= g.GetX() + g.GetWidth() && ball.GetY() >= g.GetY()  && ball.GetY() <= g.GetY() + g.GetHeight()

	if (hasCollision){
		ball.SpeedX *= -1
		ball.SpeedY *= -1
	}
}


func (ball *Ball) CheckEdges(screenWidth int, screenHeight int){

    if ball.X <= 0 || ball.X >= screenWidth {
        ball.SpeedX *= -1
    }

    if ball.Y <= 0 || ball.Y >= screenHeight {
        ball.SpeedY *= -1
    }

}
