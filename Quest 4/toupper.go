package piscine

func ToUpper(s string) string{
	lst := []rune(s)
	res := ""

	for i := 0 ; i < len(lst); i++{
		if lst[i] >= 'a' && lst[i] <= 'z'{
			res += string(lst[i] - ('a' - 'A'))
		
		} else {
			res+= string(lst[i])
		} 

	}
	return res

	
}
