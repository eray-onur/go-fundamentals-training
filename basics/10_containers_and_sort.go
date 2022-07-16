package main

import (
	"container/list"
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	// linkedLists()
	kids := []Person {
		{"Jill", 9},
		{"Jake", 11},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}

func linkedLists() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e:= x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}