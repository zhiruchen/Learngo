package main

import "fmt"

func main() {
	x := 0
	ptr_of_int := new(int)
	fmt.Println(x)

	pointer_test(&x)
	pointer_test(ptr_of_int)

	x1, y1 := 1, 2
	swap_number(&x1, &y1)
	fmt.Println(x1, y1)

}

func pointer_test(xPtr *int) {
	*xPtr = 1
	fmt.Println(*xPtr)
}

func swap_number(num1 *int, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}
