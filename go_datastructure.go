package main

import "fmt"

func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
    array_example()
    slice_example()
}

func array_example() {
    x := [5]float64{ 98, 93, 77, 82, 83}

    var total float64
    for _, value := range x {
       total += value
    }

    fmt.Println(total / float64(len(x)))

}

func slice_example() {
    // create a slice
    x := make([]float64, 4)
    x1 := make([]float64, 1, 5)

    x = append(x, 1, 2, 3)
    x1 = append(x1, 100)

    fmt.Println(x)
    fmt.Println(x1)

    slice1 := []int{4, 5, 6}
    slice2 := append(slice1, 7, 8)

    fmt.Println(slice1)
    fmt.Println(slice2)

    slice3 := []int{1, 2, 3}
    slice4 := make([]int, 2)
    copy(slice4, slice3)
    fmt.Println(slice3, slice4)
}

func map_example() {

}
