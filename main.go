package main

import (
	"fmt"
	"math/rand"
	"time"
)

func swap(slice []int, index1 int, index2 int) []int {
	slice[index1], slice[index2] = slice[index2], slice[index1]
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

func merge(left []int, right []int) []int {

	totalSize := len(left) + len(right)

	merged := make([]int, 0, totalSize)

	leftIndex := 0
	rightIndex := 0

	for leftIndex+rightIndex < totalSize {

		if leftIndex > len(left)-1 {
			rightItem := right[rightIndex]

			merged = append(merged, rightItem)
			rightIndex++

			continue
		}

		if rightIndex > len(right)-1 {
			leftItem := left[leftIndex]

			merged = append(merged, leftItem)
			leftIndex++

			continue
		}

		leftItem := left[leftIndex]
		rightItem := right[rightIndex]

		if leftItem < rightItem {
			merged = append(merged, leftItem)
			leftIndex++
		} else {
			merged = append(merged, rightItem)
			rightIndex++
		}
	}

	return merged
}

func mergeSort(toSort []int) []int {

	toMerge := make([][]int, 0, len(toSort))

	for _, item := range toSort {
		toMerge = append(toMerge, []int{item})
	}

	for len(toMerge) > 1 {
		newToMerge := make([][]int, 0)

		for i := 0; i < len(toMerge); i += 2 {
			var merged []int

			if i+1 >= len(toMerge) {
				merged = merge(toMerge[i], []int{})
			} else {
				merged = merge(toMerge[i], toMerge[i+1])
			}

			newToMerge = append(newToMerge, merged)
		}

		toMerge = newToMerge
	}

	return toMerge[0]
}

func getRandomInRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func quickSort(toSort []int, min int, max int) []int {

	pivotIndex := min

	pivot := toSort[pivotIndex]

	//swap the last item with the pivot
	toSort = swap(toSort, pivotIndex, max)

	leftIndex := min
	rightIndex := max - 1

	for leftIndex < rightIndex {
		if toSort[leftIndex] > pivot && toSort[rightIndex] < pivot {
			toSort = swap(toSort, leftIndex, rightIndex)
		}

		if toSort[leftIndex] < pivot {
			leftIndex++
		}

		if toSort[leftIndex] > pivot && toSort[rightIndex] < pivot {
			toSort = swap(toSort, leftIndex, rightIndex)
		}

		if toSort[rightIndex] > pivot {
			rightIndex--
		}
	}

	toSort = swap(toSort, leftIndex, max)

	if pivotIndex-min > 1 {
		toSort = quickSort(toSort, min, pivotIndex-1)
	}

	if max-pivotIndex > 1 {
		toSort = quickSort(toSort, pivotIndex+1, max)
	}

	return toSort
}

func main() {
	toSort := []int{5, 10, 15, 25, 0, 0, 0, 1, 1, 1, 2, 230}
	sorted := quickSort(toSort, 0, len(toSort)-1)
	fmt.Println(sorted)
}
