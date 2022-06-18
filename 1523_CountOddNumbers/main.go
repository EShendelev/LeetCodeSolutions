package main

import "fmt"

func main() {
	fmt.Println(countOdds(8, 10))
}

func countOdds(low int, high int) int {
	count := 0
	for low <= high {
		low++
		if low%2 == 0 {
			count++
		}
	}
	return count
}
