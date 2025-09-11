package piscine

func Compare(a, b string) int {
	r1 := []rune(a)
	r2 := []rune(b)

	for i := 0; i < len(r1) && i < len(r2); i++ {
		if r1[i] < r2[i] {
			return -1
		}
		if r1[i] > r2[i] {
			return 1
		}
	}

	if len(r1) < len(r2) {
		return -1
	} else if len(r1) > len(r2) {
		return 1
	}

	return 0
}