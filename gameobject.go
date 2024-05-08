package main

// Represents the core behaviour for a game object 
type GameObject interface {
	Update()
	Display() (rune)
	GetX()(int)
	GetY()(int)
}