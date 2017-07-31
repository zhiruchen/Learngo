package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

type Shape interface {
	area() float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func distance(x1, x2, y1, y2 float64) float64 {
	a := (x2 - x1)
	b := (y2 - y1)
	return math.Sqrt(a*a + b*b)
}

func (rect *Rectangle) area() float64 {
	l := distance(rect.x1, rect.y1, rect.x1, rect.y2)
	w := distance(rect.x1, rect.y1, rect.x2, rect.y1)

	return l * w
}

func main() {

	/* 因为Circle, Rectangle都实现了返回值为float64的area方法
	   所以这两个类型都实现了Shape接口
	*/

	c := Circle{x: 0, y: 0, r: 5}
	rect := Rectangle{1, 2, 10, 100}
	fmt.Println(totalArea(&c, &rect))
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}

	return area
}
