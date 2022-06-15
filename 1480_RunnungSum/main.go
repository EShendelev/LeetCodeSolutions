package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))
}

func runningSum(nums []int) []int {
	res := make([]int, 0, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = nums[i] + sum
		res = append(res, sum)
	}
	return res
}
