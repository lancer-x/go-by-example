package main

import (
	"fmt"
	"math"
)

type point struct {
	X int
	Y int
}

type insideDot struct {
	dotX int
	dotY int

}

type circle struct {
	Radius float64
	Center point
	insideDot
}

func (c circle) calcArea() float64 {
	return float64(math.Pi * c.Radius * c.Radius)
}

func (c *circle) calcArea2() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (p *point) printPoint() {
	fmt.Printf("%d,,,%d\n", p.X, p.Y)
}

func (d *insideDot) showDot() {
	fmt.Printf("%d----%d\n", d.dotX, d.dotY)
}



func main()  {
	circlePtr := newCirclePtr()
	circle1 := newCircle()
	fmt.Println(circle1.calcArea())
	fmt.Println(circle1.calcArea2())

	fmt.Println(circlePtr.calcArea())
	fmt.Println(circlePtr.calcArea2())

	circlePtr.Center.printPoint()
	circle1.Center.printPoint()

	circle1.showDot()
	(*circlePtr).showDot()
}

func newCirclePtr()  *circle{
	return &circle{
		Center:point{X:10, Y:20},
		Radius:15,
		insideDot: insideDot{
			dotX: 11,
			dotY: 22,
		},
	}
}

func newCircle() circle {
	return circle{
		Center:point{X:10, Y:20},
		Radius:15,
		insideDot: insideDot{
			dotX: 22,
			dotY: 33,
		},
	}
}

