package piscine

func IsPrintable(s string) bool {
	lst := []rune(s)

	for i := 0; i < len(lst); i++{
		if lst[i] < 32 || lst[i] > 126 {
			return false
		}
	}

	return true
}