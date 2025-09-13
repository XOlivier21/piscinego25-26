package piscine

func BasicJoin(elems []string ) string {
	res := ""

	for i := 0; i < len(elems); i++{
		res += string(elems[i])
	}

	return res
}