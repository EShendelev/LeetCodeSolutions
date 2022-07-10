package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(arr))
}

func maxSubArray(nums []int) int {
	currentSum := 0
	maxSum := math.MinInt
	for i := range nums {
		currentSum += nums[i]
		if currentSum < nums[i] {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
