package piscine

func IterativeFactorial(n int) int{
	if n < 0 {
		return 0
	}
	resultat := 1
	for i := 1; i <= n; i++{
		resultat *= i
	

	}
	return resultat
}

