package main

import "fmt"

func romanToInteger(s string) int {
	prev := 0
	integer := 0
	romanMap := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(s[i])
		curr := romanMap[s[i]]
		if curr < prev {
			integer -= curr
		} else {
			integer += curr
		}
		prev = curr

	}
	return integer
}

func main() {
	fmt.Println(romanToInteger("MCMXCIV"))
}
