package piscine

func FindNextPrime(n int) int {
	if n <= 2 {
		return 2
	}

	for {
		prime := true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				prime = false
				break
			}
		}
		if prime {
			return n
		}
		n++
	}
}