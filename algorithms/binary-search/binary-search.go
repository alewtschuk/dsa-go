package main

import "fmt"

// Binary search is a search algorithm that works by finding the
// position of a target value within a sorted array. It repeatedly divides
// the search interval in half till the target value is found or the
// search space is empty. Returns -1 if search space is exhausted
//
// The algorithm step by step functions as follows:
//   - Divide the search space into two halves using the midpoint
//   - Compare the middle element of the search space with the target
//   - If the target is found function terminates and returns
//   - If the target is not found at the midpoint choose which half will be searched next
//   - If the target is larget than the midpoint it will search the right side first
//   - If the target is smaller than the midpoint it will search the left side first
//   - This is continued till the target is found or the search space is exhausted.
//
// NOTE: The input array must be sorted. This algorithm for simplicity assumes sorted input.
func binarySearch(arr []int, target int, low int, high int) int {
	mid := (low + high) / 2

	// Base case where midpoint is the target
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target { // The value of mid is less than target, right side is searched next

		// Set new high to index one to the left of old midpoint
		low = mid + 1
		// Recursively call binary search with the new array from 0 to the new high
		return binarySearch(arr, target, low, high)
	} else {
		// Mid is larger than target right side is searched next
		if arr[mid] > target {
			// Set new low to index one to the right of the old midpoint
			high = mid - 1
			// Recursively call binary search with the new array from the new low to the high
			return binarySearch(arr, target, low, high)
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 4, 6, 9, 10, 14, 17, 28, 34, 46, 55, 69, 78, 85, 90, 99, 100}
	target := 2
	result := binarySearch(arr, target, 0, len(arr)-1)
	if result != -1 {
		fmt.Printf("Target %d found at index %d\n", target, result)
	} else {
		fmt.Println("Search space has been exhausted")
	}
}
