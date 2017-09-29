//exercise 10.3

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Polar struct {
	point *Point
	z     float64
}

func (p *Point) Abs() float64 {
	a := p.x * p.y
	return math.Abs(a)
}

func (p *Polar) Scale() float64 {
	sc := p.point.Abs() * p.z
	return sc
}

func main() {
	point := &Point{3, 5}
	var polar Polar
	polar.z = 0.222222
	polar.point = point
	fmt.Println(polar.Scale())
}
