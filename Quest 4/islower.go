package piscine

func IsLower(s string) bool {

	lst := []rune(s)

	for i := 0; i < len(lst); i++ {
		if lst[i] < 'a' || lst[i] > 'z' {
			return false
		}
	}
	return true
}