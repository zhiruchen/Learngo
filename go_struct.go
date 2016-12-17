package main

import "fmt"
import "math"

// type: introduces a new type
// struct: defining a struct type
type Circle struct {
    x, y, r float64
}

type Rectangle struct {
    x1, x2, y1, y2 float64
}

type Person struct {
    Name string
}

type Android struct {
    Person
    Model string
}


func main() {
    c := Circle{x: 0, y: 0, r: 5}
    fmt.Println(c.x, c.y, c.r)
    fmt.Println(circleArea(c))


    c1 := new(Circle)
    fmt.Println(c1.x, c1.y, c1.r)

    c2 := Circle{0, 0, 100}
    circle_area := c2.area()
    fmt.Println(circle_area)

    rect := Rectangle{1, 2, 10, 100}
    fmt.Println(rect.area())

    // go embedded type
    a := new(Android)
    a.Talk()

    b := Android{
        Person: Person{
            Name: "zhiru",
        },
        Model: "典范",
    }
    b.Talk()

}

func circleArea(c Circle) float64 {
    return math.Pi * c.r * c.r
}

// Circle method
func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func distance(x1, x2, y1, y2 float64) float64 {
    a := (x2 - x1)
    b := (y2 - y1)
    return math.Sqrt(a * a + b * b)
}

func (rect *Rectangle) area() float64 {
    l := distance(rect.x1, rect.y1, rect.x1, rect.y2)
    w := distance(rect.x1, rect.y1, rect.x2, rect.y1)

    return l * w
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is: ", p.Name)
}