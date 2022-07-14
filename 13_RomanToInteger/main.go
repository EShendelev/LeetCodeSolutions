package main

import "fmt"

func main() {
	fmt.Println(romanToInt(("LVIII")))
}

func romanToInt(s string) int {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := 0; i < len(s)-1; i++ {
		if m[rune(s[i])] < m[rune(s[i+1])] {
			res -= m[rune(s[i])]
		} else {
			res += m[rune(s[i])]
		}
	}
	res += m[rune(s[len(s)-1])]
	return res
}
