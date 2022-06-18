package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(10))
}

func subtractProductAndSum(n int) int {
	if n < 10 {
		return 0
	}
	arr := []int{}
	sum, prod := 0, 1
	for i := 0; n >= 1; i++ {
		arr = append(arr, n%10)
		// fmt.Println(arr)
		n = n / 10
		sum += arr[i]
		prod *= arr[i]
	}

	return prod - sum
}
