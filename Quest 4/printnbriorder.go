package piscine

func PrintNbrInOrder(n int) int {
	lst := []int{}
	for n > 0 {
		lst = append(lst, n%10)
		n = n / 10
	}

	for i := 0; i < len(lst); i++ {
		for j := i+1 ; j < len(lst); j++ {
			if lst[i] > lst[j]{
				lst[j] , lst[i]  = lst[i] , lst[j]
			}
		}

	}
	num := 0
	for i := 0; i < len(lst); i++ {
		num = num*10 + lst[i]
	}
	return num
}