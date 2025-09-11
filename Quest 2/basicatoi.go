package piscine

func BasicAtoi(s string) int{

	resultat := 0

	for _ , r := range s {
		if r >= '0' && r <= '9' {
			resultat = resultat*10 + int(r - '0')
	
		}

	}
	return resultat
}