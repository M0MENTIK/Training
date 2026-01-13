package main

type Shape interface {
	Area() float64
}

//

//

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//

//

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
