package piscine

func IsNumeric(s string) bool {
	lst := []rune(s)
	for i := 0; i < len(lst); i++{
		if lst[i] < '0' || lst[i] > '9'{
			return false
		}


	}
	return true

}