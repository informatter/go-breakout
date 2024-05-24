package main

// Represents the core behaviour for a game object 
type GameObject interface {
	Update(engine Engine)
	Display(engine *Engine)
	GetX()(int)
	GetY()(int)
	GetHeight()(int)
	GetWidth()(int)
	// Checks if the game object is within the screen boundaries
	CheckEdges(screenWidth int, screenHeight int)
	CheckCollision(gameObject GameObject)
	GetActiveState() (bool)
}