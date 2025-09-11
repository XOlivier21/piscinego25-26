package piscine

func IsAlpha(s string) bool{
	lst := []rune(s)

	for i := 0; i < len(lst); i++{
		if (lst[i] < 'a' || lst[i] > 'z') && (lst[i] < 'A' || lst[i] > 'Z') && (lst[i] < '0' || lst[i] > '9'){
			return false
		}
		

	}

	return true
}