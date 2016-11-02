// 打印重复行
package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	counts := make(map[string]int)  // make构造一个空字典
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}

	}


}
