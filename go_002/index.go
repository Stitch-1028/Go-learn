package main

import (
	"fmt"
	"main/go_001"
)

func main() {
	list := []int{1,2,3,4,5}
	index := go_001.ListFindIndex(list,3)
	fmt.Println(index)
}