package piscine

func IsUpper(s string) bool {

	lst := []rune(s)

	for i := 0; i < len(lst); i++ {
		if lst[i] < 'A' || lst[i] > 'Z' {
			return false
		}
	}
	return true
}