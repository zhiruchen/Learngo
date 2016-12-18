package main

import "fmt"
import mpkg "./pkg_test"  // mpkg is the alias of package pkg_test

func main() {
	xs := []float64{1, 2, 3, 4, 5, 6}
	fmt.Println(mpkg.Average(xs))
}


