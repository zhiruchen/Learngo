package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicateVowel bool = true
	removeVowel    bool = false
)

func readBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		if readBool() {
			var vI = -1
			for i, char := range word {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if readBool() {
						vI = i
					}
				}
			}
			if vI > 0 {
				switch readBool() {
				case duplicateVowel:
					word = append(word[:vI+1], word[vI:]...)
				case removeVowel:
					word = append(word[:vI], word[vI+1:]...)
				}
			}
		}
		fmt.Println(string(word))
	}
}
