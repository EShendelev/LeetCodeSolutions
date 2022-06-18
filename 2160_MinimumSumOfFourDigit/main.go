package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumSum(1254))
}

func minimumSum(num int) int {
	arr := []int{}
	for num >= 1 {
		arr = append(arr, num%10)
		num = num / 10
	}
	sort.Ints(arr)
	// fmt.Println(arr)
	return arr[0]*10 + arr[1]*10 + arr[2] + arr[3]
}
