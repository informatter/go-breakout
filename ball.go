
package main
import (
	"github.com/gdamore/tcell/v2"
	"time"
	"os"
)

// Ball entity which implements the GameObject interface
type Ball struct{
	X int
	Y int
	SpeedX int
	SpeedY int
	Active bool
	Screen tcell.Screen
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

func (ball Ball) GetActiveState() bool{
	return ball.Active
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

func resetPosition(ball *Ball){
	width,height := ball.Screen.Size()
	ball.X = width /2
	ball.Y = height / 2
}



func (ball *Ball) Update(engine Engine){

	if(ball.Active){
		ball.X+=ball.SpeedX
		ball.Y+=ball.SpeedY
	}else{

		// makes ball active again after a small pause
		engine.Player.Life-=1
		if (engine.Player.Life == 0){
			os.Exit(0)
		}
		time.Sleep(2 * time.Second)
		ball.Active = true
	}
}


func (ball *Ball)  CheckCollision(g GameObject){

	if(!ball.Active){
		return
	}

	_, isBlock := g.(*Block)

	hasCollision := 
		ball.GetX() >= g.GetX() && 
		ball.GetX() <= g.GetX() + g.GetWidth() && 
		ball.GetY() >= g.GetY() && 
		ball.GetY() <= g.GetY() + g.GetHeight()
	
	if (hasCollision  && !isBlock){
		reflect(ball,g)
		//ball.Screen.Beep()
	}
	if (hasCollision && isBlock && g.GetActiveState()){
		reflect(ball,g)
		//ball.Screen.Beep()
	}
}


func (ball *Ball) CheckEdges(screenWidth int, screenHeight int){

	if(!ball.Active){
		return
	}

    if ball.X <= 0 || ball.X >= screenWidth {
        ball.SpeedX *= -1
    }


	if (ball.Y >= screenHeight) {
		ball.Active = false
		resetPosition(ball)

    }else if (ball.Y <= 0){
		ball.SpeedY *= -1
	}

}

func reflect (ball *Ball, g GameObject){
	
	// Determine which side of the block the collision occurred
	if ball.GetX() <= g.GetX() || ball.GetX() >= g.GetX()+g.GetWidth() {
		ball.SpeedX *= -1 //Reflect X direction
	}

	if ball.GetY() <= g.GetY() || ball.GetY() >= g.GetY()+g.GetHeight() {
		ball.SpeedY *= -1 //Reflect Y direction
	}
}
