package main

import "fmt"

func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
    array_example()
    slice_example()
    map_example()
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
    elements := make(map[string]string)
    elements["haha"] = "You a Man"
    elements["heihei"] = "I am a Programmer"
    elements["code"] = "Show me the Code"

    fmt.Println(elements["code"])
    if name, ok := elements["heihei"]; ok {
        fmt.Println(name, ok)
    }

    elements1 := map[string]string {
        "a": "abc",
        "b": "bcd",
        "e": "counting star",
    }

    fmt.Println(elements1["a"])

    elements2 := map[string]map[string]string {
        "A": map[string]string{
            "name": "A",
            "position": "1",
        },
        "B": map[string]string{
            "name": "B",
            "position": "2",
        },
    }

    fmt.Println(elements2["A"]["name"], elements2["A"]["position"])
}
