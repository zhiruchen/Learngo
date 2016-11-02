package main


import (
	"fmt"
	"os"
	"strings"
)


func main() {
	s, seq := "", ""
	for _, arg := range os.Args {
		s += seq + arg
		seq = " "
	}

	fmt.Println(s)

	fmt.Println(strings.Join(os.Args, " "))
}
