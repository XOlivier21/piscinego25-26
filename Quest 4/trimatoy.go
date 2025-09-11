package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	lst := []rune(s)

	for i := 0; i < len(lst); i++ {
		if lst[i] == '-' {
			sign = -1
		}
		if lst[i] >= '0' && lst[i] <= '9' {
			result = result*10 + int(lst[i] -'0')
		}
	}

	return result * sign
}

