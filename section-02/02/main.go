// binary search
package main

import (
	"fmt"
)

func main() {
	data := []int{1, 8, 15, 17, 21, 32, 37, 47, 56, 70, 71, 74, 76, 80, 85, 97, 99}
	search := 71
	index := recursiveBinarySearch(data, search, 0, len(data))
	fmt.Printf("Found %d by recursive binary search at index %d\n", search, index)
	index = iterativeBinarySearch(data, search, 0, len(data))
	fmt.Printf("Found %d by iterative binary search at index %d\n", search, index)
}

func recursiveBinarySearch(data []int, val, startIndex, stopIndex int) int {
	if startIndex > stopIndex {
		return -1
	}
	mid := int((startIndex + stopIndex) / 2)
	if data[mid] > val {
		return recursiveBinarySearch(data, val, startIndex, mid)
	} else if data[mid] < val {
		return recursiveBinarySearch(data, val, mid, stopIndex)
	} else {
		return mid
	}
}

func iterativeBinarySearch(data []int, val, startIndex, stopIndex int) int {
	var mid int
	for startIndex <= stopIndex {
		mid = int((startIndex + stopIndex) / 2)
		if data[mid] > val {
			stopIndex = mid - 1
		} else if data[mid] < val {
			startIndex = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
