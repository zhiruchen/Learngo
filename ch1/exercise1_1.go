package main

import (
	"fmt"
	"os"
)

func main() {
	// s, seq := "", ""
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}
}
