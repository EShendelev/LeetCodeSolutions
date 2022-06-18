package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{8000, 9000, 2000, 3000, 6000, 1000}
	fmt.Println(average(nums))
}

func average(salary []int) float64 {
	sort.Ints(salary)
	sum := 0
	for i := 1; i < len(salary)-1; i++ {
		sum += salary[i]
	}
	return float64(sum) / float64(len(salary)-2)
}
