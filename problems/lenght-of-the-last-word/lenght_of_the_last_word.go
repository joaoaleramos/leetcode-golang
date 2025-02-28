package main

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
