package main

import "fmt"

type Rectangle struct {
	length, width float64
}

// method with value receiver
func (r Rectangle) Area() float64 {
	return r.width * r.length
}

// method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {
	rect := Rectangle{
		length: 20,
		width:  10,
	}
	area := rect.Area()
	fmt.Println(area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println(area)
}
