package main

import (
	"fmt"
	"math"
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

func getMergeSliceSize(length int) int {
	return int(math.Ceil(float64(length) / 2))
}

func mergeSort(toSort []int) []int {

	halfLength := int(getMergeSliceSize(len(toSort)))
	sorted := make([][]int, halfLength)

	sortedIndex := -1

	for index, item := range toSort {
		var newSlice []int

		if index%2 == 0 {
			slice := make([]int, 0, 2)
			newSlice = append(slice, item)

			sortedIndex++
		} else {
			slice := sorted[sortedIndex]

			if slice[0] > item {
				newSlice = []int{item, slice[0]}
			} else {
				newSlice = append(slice, item)
			}
		}

		sorted[sortedIndex] = newSlice
	}

	fmt.Println(sorted)
	return sorted[0]
}

func main() {
	toSort := []int{24, 13, 9, 64, 7, 23, 34, 47, 5}
	mergeSort(toSort)
	//fmt.Println(sorted)
}
