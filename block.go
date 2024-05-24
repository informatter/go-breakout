package main

import (
	"github.com/gdamore/tcell/v2"
	"strings"
)

type Block struct {
	Width  int
	Height int
	X      int
	Y      int
	Style tcell.Style
	Active bool
	ScoreComputed bool
}

func (block Block) GetX() int{
	return block.X
}

func (block Block) GetY() int{
	return block.Y
}

func (block Block) GetHeight() int{
	return block.Height
}

func (block Block) GetWidth() int{
	return block.Width
}

func (block Block) GetActiveState() bool{
	return block.Active
}

func (block Block) Display(engine *Engine) {
	shape := strings.Repeat(" ", block.Width)
	// Style := tcell.StyleDefault.Background(tcell.ColorPapayaWhip).Foreground(tcell.ColorAntiqueWhite)

	renderGameObject(
		engine.Screen,
		block.GetX(),
		block.GetY(),
		block.GetX() + block.Width,
		block.GetY() + block.Height,
		block.Style,
		shape,
	)
}

func (block *Block) Update(engine Engine){
	if (!block.Active && !block.ScoreComputed){
		block.ScoreComputed = true
		engine.Player.Score+=1
	}
}


func (block *Block) CheckEdges(screenWidth int, screenHeight int){

}

func (block *Block) CheckCollision(g GameObject){
	_, ok := g.(*Ball)
	if !ok {
		return
	}

	hasCollision := 
		g.GetX() >= block.GetX() && 
		g.GetX() <= block.GetX() + block.GetWidth() &&
		g.GetY() >= block.GetY()  && 
		g.GetY() <= block.GetY() + block.GetHeight()

	if (hasCollision){
		block.Style = tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
		block.Active = false
	}

}