package piscine

func Swap(a *int, b *int) {
	var nb int
	nb = *a
	*a = *b
	*b = nb
	
}