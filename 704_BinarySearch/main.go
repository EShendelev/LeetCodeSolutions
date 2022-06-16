package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
}

func search(nums []int, target int) int {
	sort.Ints(nums)
	minInd := 0
	maxInd := len(nums) - 1
	for minInd <= maxInd {
		mid := (minInd + maxInd) / 2
		guess := nums[mid]

		if guess == target {
			return mid
		} else if guess > target {
			maxInd = mid - 1
		} else {
			minInd = mid + 1
		}
	}
	return -1
}
