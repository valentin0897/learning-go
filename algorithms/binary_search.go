package main

import "fmt"

func main() {
	x := []int{-2, -1, 0, 1, 7, 50, 60, 400}
	fmt.Println(x)
	index := binary_search(x, -2)
	fmt.Println(index)
}

func binary_search(list []int, target int) int {
	var mid int
	low := 0
	high := len(list) - 1
	result := -1

	for low <= high {
		mid = (low + high) / 2
		if list[mid] == target {
			result = mid
			break
		} else if list[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}
