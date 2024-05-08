
package main

type Ball struct{
	X int
	Y int
	SpeedX int
	SpeedY int
}

func (h Ball) GetX() int{
	return h.X
}

func (h Ball) GetY() int{
	return h.Y
}

func (h Ball) Display() rune{
	return '\u26AA'
}

func (h *Ball) Update(){

	h.X+=h.SpeedX
	h.Y+=h.SpeedY
}
