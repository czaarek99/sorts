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

func insertionSort(toSort []int) []int {

	sorted := make([]int, 0, len(toSort))

	for _, item := range toSort {
		sorted = append(sorted, item)

		for i := len(sorted) - 1; i > 0; i-- {
			prevIndex := i - 1

			if sorted[i] < sorted[prevIndex] {
				swap(sorted, i, prevIndex)
			} else {
				break
			}

		}
	}

	return sorted
}

func main() {
	toSort := []int{24, 13, 9, 64, 7, 23, 34, 47}
	sorted := insertionSort(toSort)
	fmt.Println(sorted)
}
