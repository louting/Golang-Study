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
	/*
		此处不能使用 var polar *Polar 这样的定义方式，不然下面的语句
		polar.z 属于一个指针，应该不能为指针赋值
		报错为：runtime error: invalid memory address or nil pointer dereference
	*/
	polar.z = 10
	polar.point = point
	fmt.Println(polar.Scale())
}

/* result
150
*/
