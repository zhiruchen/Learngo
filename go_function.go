package main

import "fmt"

func main() {
    xs := []float64{ 11.22, 22.11, 88.66, 100.001 }
    fmt.Println(average(xs))

    a, b := get_multi_return_value()
    fmt.Println(a, b)

    sum := sum_of_multi_args(1, 2, 3)
    fmt.Println(sum)

    xs1 := []int{4, 5, 6}
    fmt.Println(sum_of_multi_args(xs1...))

    add := func(x, y int) int {
        return x + y
    }

    fmt.Println(add(1, 2))

    nextEven := makeEvenGenerator()
    fmt.Println(nextEven())
    fmt.Println(nextEven())
    fmt.Println(nextEven())

    fmt.Println(fact(10))

    // defer_panic_test()
    defer_panic_test1()

}

func average(xs []float64) float64 {
    total := 0.0
    for _, value := range xs {
        total += value
    }

    return total / float64(len(xs))
}

func get_multi_return_value() (int, int) {
    return 5, 6
}

func sum_of_multi_args(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }

    return total
}

func makeEvenGenerator() func() uint {
    i := uint(0)

    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func fact(x uint) uint {
    if x == 0 {
        return 1
    }

    return x * fact(x - 1)
}

func defer_panic_test() {
    panic("PANIC")  // 停止对defer_panic_test的调用
    str := recover()
    fmt.Println(str)
}

func defer_panic_test1() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}