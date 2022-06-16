package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 2))
}

func searchInsert(nums []int, target int) int {
	minInd := 0
	maxInd := len(nums) - 1

	for minInd <= maxInd {
		mid := minInd + (maxInd-minInd)/2
		guess := nums[mid]

		if guess == target {
			return mid
		} else if guess > target {
			maxInd = mid - 1
		} else {
			minInd = mid + 1
		}
	}
	return minInd
}
