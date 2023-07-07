package main

import (
	"fmt"
)

func main() {
	fmt.Println("Output: ", search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func binarySearch(arr []int, left int, right int, target int) int {
	if right >= left {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			return binarySearch(arr, left, mid-1, target)
		}

		return binarySearch(arr, mid+1, right, target)
	}

	return -1
}
