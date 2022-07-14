package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {

		val := target - nums[i]
		if _, ok := m[val]; ok {
			return []int{i, m[val]}
		}
		m[nums[i]] = i
	}
	return nil
}
