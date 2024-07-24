package main

import (
	"fmt"
	"fungolang/sort"
)

func main() {
	list := []int{3, 92, 10, 5, 44, 32, 6, 3, 7, 11, 12, 17, 2, 54, 36, 23, 14}

	listCopy := make([]int, len(list))
	copy(listCopy, list)

	bubbleSortedList := sort.BubbleSort(listCopy)
	fmt.Println("Bubble Sort Example:", list, "->", bubbleSortedList)
}