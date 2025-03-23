package main

import (
	"fmt"
	"os"
	"strconv"
)

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
	// Calculate the midpoint in an integer overflow safe way: low value + (find the length of the array) / 2
	mid := low + (high-low)/2

	// Base case target is found
	if arr[mid] == target {
		return mid
	} else if low == high {
		// Secondary case where low and high are equivalent and search space is exhausted. Target is not found
		return -1
	} else if arr[mid] < target { // The value of mid is less than target, right side is searched next
		// Recursively call binary search with the new array from the new low (mid+1) to the high
		return binarySearch(arr, target, mid+1, high)
	} else {
		// Mid is larger than target right side is searched next
		if arr[mid] > target {
			// Recursively call binary search with the new array from low to the new high (mid+1)
			return binarySearch(arr, target, low, mid-1)
		}
	}

	return -1
}

func main() {
	arr := []int{2, 9, 16, 23, 30, 37, 44, 51, 58, 65, 72, 79, 86, 93, 100, 107, 114, 121, 128, 135, 142, 149, 156, 163, 170, 177, 184, 191, 198, 205, 212, 219, 226, 233, 240, 247, 254, 261, 268, 275, 282, 289, 296, 303, 310, 317, 324, 331, 338, 345, 352, 359, 366, 373, 380, 387, 394, 401, 408, 415, 422, 429, 436, 443, 450, 457, 464, 471, 478, 485, 492, 499, 506, 513, 520, 527, 534, 541, 548, 555, 562, 569, 576, 583, 590, 597, 604, 611, 618, 625, 632, 639, 646, 653, 660, 667, 674, 681, 688, 695, 702, 709, 716, 723, 730, 737, 744, 751, 758, 765, 772, 779, 786, 793, 800, 807, 814, 821, 828, 835, 842, 849, 856, 863, 870, 877, 884, 891, 898, 905, 912, 919, 926, 933, 940, 947, 954, 961, 968, 975, 982, 989, 996}
	args := os.Args
	for _, arg := range args[1:] {
		arg, _ := strconv.Atoi(arg)
		result := binarySearch(arr, arg, 0, len(arr)-1)
		if result != -1 {
			fmt.Printf("Target %d found at index %d\n", arg, result)
		} else {
			fmt.Println("Search space has been exhausted")
		}
	}
}
