package main

type Location struct {
	X, Y int
}

func (l Location) Shift(shiftX, shiftY int) *Location {
	nextX := l.X + shiftX
	nextY := l.Y + shiftY
	return &Location{nextX, nextY}
}

func NewLocation(x, y int) *Location {
	return &Location{x, y}
}
