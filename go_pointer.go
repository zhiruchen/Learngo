package main

import "fmt"

func main() {
    x := 0
    ptr_of_int := new(int)
    fmt.Println(x)

    pointer_test(&x)
    pointer_test(ptr_of_int)

}

func pointer_test(xPtr *int) {
    *xPtr = 1
    fmt.Println(*xPtr)
}