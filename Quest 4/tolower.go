package piscine

func ToLower(s string) string{
	lst := []rune(s)

	res := ""

	for i := 0; i < len(lst); i++{
		if lst[i] >= 'A' && lst[i] <= 'Z'{
			res += string(lst[i] + 'a'-'A')

		} else{
			res += string(lst[i])
			
		}


	}
	return res
}