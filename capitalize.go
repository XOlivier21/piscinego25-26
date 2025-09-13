package piscine

func Capitalize(s string) string {
	lst := []rune(s)

	if len(lst) > 0 && lst[0] >= 'a' && lst[0] <= 'z' {
		lst[0] = lst[0] - ('a' - 'A')
	}

	for i := 1; i < len(lst)-1; i++ {
		if lst[i] == '!' || lst[i] == '?' || lst[i] == '.' || lst[i] == ' ' || lst[i] == '+' || lst[i] == '-' {
			if lst[i+1] >= 'a' && lst[i+1] <= 'z' {
				lst[i+1] = lst[i+1] - ('a' - 'A')
				continue
			}
			if lst[i+1] == ' ' {
				lst[i+2] = lst[i+2] + ('a' - 'A')
			}
		}

		if lst[i+1] >= 'A' && lst[i+1] <= 'Z' {
			lst[i+1] = lst[i+1] + ('a' - 'A')
		}
	}

	return string(lst)
}