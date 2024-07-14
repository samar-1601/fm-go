package main

const PI = 3.14

type Circle struct {
	radius float64
}

func (c Circle) circumference() {
	println(2 * PI * c.radius)
}

func main() {
	circle := Circle{
		radius: 10,
	}
	circle.circumference()
}
