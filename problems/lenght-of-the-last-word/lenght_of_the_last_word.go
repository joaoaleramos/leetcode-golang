package main

import "strings"

func lengthOfLastWord(s string) int {
	var length int

	if len(s) == 0 {
		return 0
	} else {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == ' ' {
				if length == 0 {
					continue
				}
				break
			}
			length++
		}
	}

	return length

}

func lengthOfLastWordOptimized(s string) int {
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}
	lastWord := words[len(words)-1]
	return len(lastWord)
}
