package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	rotate(arr, 3)
	fmt.Println(arr)
	// fmt.Println(2 % 1)

}

func rotate(nums []int, k int) {
	l := len(nums)
	// arr := nums
	copy(nums, append(nums[l-k%l:], nums[:l-k%l]...))

}
