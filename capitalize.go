package piscine

func Capitalize(s string)  string{

	lst := []rune(s)

	res := ""
	if lst[0] >= 'a' && lst[0] <='z'{
		lst[0] = lst[0] - ('a' - 'A')

	}
	for i := 1; i < len(lst); i++{
		if lst[i] >= 'A' && lst[i] <= 'Z'{
			lst[i] = lst[i] + ('a' - 'A')


	}

	
}