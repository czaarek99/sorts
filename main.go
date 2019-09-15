package main

import (
	"fmt"
)

func swap(slice []int, index1 int, index2 int) []int {
	firstVal := slice[index1]
	secondVal := slice[index2]

	slice[index1] = secondVal
	slice[index2] = firstVal

	return slice
}

func bubbleSort(toSort []int) []int {

	for i := len(toSort) - 1; i != 0; i-- {
		for j := 0; j < i; j++ {
			nextIndex := j + 1

			if nextIndex < len(toSort) && toSort[j] > toSort[nextIndex] {
				toSort = swap(toSort, j, nextIndex)
			}
		}
	}

	return toSort
}

func main() {
	toSort := []int{5, 9, 2, 7, 1}
	sorted := bubbleSort(toSort)
	fmt.Println(sorted)
}
