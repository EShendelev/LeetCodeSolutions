package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(arr))
}

func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)
	return nums
}
