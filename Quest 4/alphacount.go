package piscine

func AlphaCount(s string) int {
	lst := []rune(s)
	count := 0

	for i:= 0; i < len(lst); i++{
		if (lst[i] >= 'a' && lst[i] <= 'z') || (lst[i] >= 'A' && lst[i] <= 'Z'){
			count++
	}
}
	return count
}
