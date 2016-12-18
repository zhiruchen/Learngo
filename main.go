package main

import "fmt"
import mpkg "./mmath"  // mpkg is the alias of package mmath

func main() {
	xs := []float64{1, 2, 3, 4, 5, 6}
	fmt.Println(mpkg.Average(xs))
}


