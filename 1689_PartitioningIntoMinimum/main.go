package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(minPartitions("27346209830709182346"))
}

func minPartitions(n string) int {
	arr := strings.Split(n, "")
	var max int
	for i := range arr {
		digit, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Errorf("Not integer")
		}
		if digit > max {
			max = digit
		}
	}
	return max
}
