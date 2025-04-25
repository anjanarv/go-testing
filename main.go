package main

import (
	"fmt"
)

func selectionSort(n []int) []int {
	// find the smallest in the list iteratively
	// swap it for the place of first unsorted item
	// find the next smallest in the remaining number of items

	for i := 0; i < len(n)-1; i++ {
		smallestIndex := i
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[smallestIndex] {
				smallestIndex = j // find the smallest
			}
		}
		n[smallestIndex], n[i] = n[i], n[smallestIndex] // swapping the smallest
	}

	fmt.Println("Selection Sorted Array: ", n)

	return n
}
