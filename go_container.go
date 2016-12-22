package main

import (
	"fmt"
	"sort"
	"container/list"
)

type Person struct {
	Name string
	Age int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value.(int))
	}

	persons := []Person{
		{"Jack", 24},
		{"John", 26},
	}
	sort.Sort(ByName(persons))
	fmt.Println(persons)
}
