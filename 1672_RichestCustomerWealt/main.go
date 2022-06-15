package main

import "fmt"

func main() {
	arr := [][]int{{6, 59, 64, 19, 30, 76, 71, 86, 90, 25, 56, 17, 19, 72, 61, 56, 24, 40, 35, 39, 67, 28, 52, 11, 82, 72, 8, 82, 81, 47}}
	fmt.Println(maximumWealth(arr))
	fmt.Println(len(arr))
}

func maximumWealth(accounts [][]int) int {
	max := 0
	sum := 0
	if len(accounts) > 1 {
		for i := range accounts {
			if sum > max {
				max = sum
			}
			sum = 0
			for _, v := range accounts[i] {
				sum += v
			}
		}
	} else {
		for _, v := range accounts[0] {
			max += v
		}
	}
	return max
}
