package main

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle)  Area() float64 {
	return 0
}

func (c Circle) Area() float64  {
	return 0
}