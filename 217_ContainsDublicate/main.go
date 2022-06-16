package main

import "fmt"

func main() {
	nums := []int{0, 0}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	hash := make(map[int]int, len(nums))

	for _, v := range nums {
		hash[v]++
		if hash[v] > 1 {
			return true
		}
	}
	return false
}
