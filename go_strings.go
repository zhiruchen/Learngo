package main

import (
	"fmt"
	"strings"
)

/*
$ go run go_strings.go
true 2 true true 1 a-b aaaaa bbaa [a b c d e] test TEST

 */
func main() {
	fmt.Println(
		strings.Contains("test", "es"),  // true
		strings.Count("test", "t"),      // 2
		strings.HasPrefix("test", "te"),  // true
		strings.HasSuffix("test", "st"),  // true
		strings.Index("test", "es"),      // 1
		strings.Join([]string{"a", "b"}, "-"),  // a-b
		strings.Repeat("a", 5),  // aaaaa
		strings.Replace("aaaa", "a", "b", 2),  // bbaa
		strings.Split("a b c d e", " "), // []string{a, b, c, d, e}
		strings.ToLower("TEST"),  // test
		strings.ToUpper("test"),  // TEST
	)
}