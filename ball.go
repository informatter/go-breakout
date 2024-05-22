
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

func (ball Ball) Display() rune{
	return '\u26AA'
}

func (ball *Ball) Update(){

	ball.X+=ball.SpeedX
	ball.Y+=ball.SpeedY

}


func (ball *Ball) CheckEdges(screenWidth int, screenHeight int){

    if ball.X <= 0 || ball.X >= screenWidth {
        ball.SpeedX *= -1
    }

    if ball.Y <= 0 || ball.Y >= screenHeight {
        ball.SpeedY *= -1
    }

}
