package main

import (
	"fmt"
	"strings"
)

func main() {
	ip := "255.100.50.0"
	fmt.Println(defangIPaddr(ip))
}

func defangIPaddr(address string) string {
	ip := strings.Split(address, ".")
	return strings.Join(ip, "[.]")
}
