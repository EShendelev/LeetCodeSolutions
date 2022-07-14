package main

import "fmt"

func main() {
	n := 15
	for i := range fizzBuzz(n) {
		fmt.Println(fizzBuzz(n)[i])
	}
}

func fizzBuzz(n int) []string {
	arr := []string{}
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			arr = append(arr, "FizzBuzz")
		case i%5 == 0:
			arr = append(arr, "Buzz")
		case i%3 == 0:
			arr = append(arr, "Fizz")
		default:
			arr = append(arr, fmt.Sprint(i))
		}
	}
	return arr
}
