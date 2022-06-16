package main

import "fmt"

func main() {
	fmt.Println(guessNumber(5))
}

func guessNumber(n int) int {
	left := 1
	right := n

	for left < right {
		num := left + (right-left)/2
		guess := 0 //guess(num)

		if guess == 0 {
			return num
		} else if guess == 1 { // num < pick
			left = num + 1
		} else {
			right = num - 1
		}
	}
	return left
}
