package main

func addBinary(a string, b string) string {
	carry := 0
	result := make([]byte, 0, max(len(a), len(b))+1)
	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		carry = sum / 2
		result = append(result, byte(sum%2+'0'))
	}
	for k, n := 0, len(result); k < n/2; k++ {
		result[k], result[n-1-k] = result[n-1-k], result[k]
	}
	return string(result)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
