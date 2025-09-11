package piscine

func Index(a,b string) int {
	t := []rune(a)
	m := []rune(b)

	if len(m) == 0 {
		return 0
	}

	for i := 0; i < len(t) - len(m); i++{
		match := true
		for j := 0; j < len(m); j++ {
			if t[i+j] != m[j] {
				match = false
				break
		}
		if match {
			return i
		}



	}



	
}
return -1

}