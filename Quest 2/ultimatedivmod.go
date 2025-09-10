package piscine

func UltimateDivMod(a *int, b *int){

	quotient := *a / *b
	reste := *a % *b

	*a = quotient
	*b = reste
	
}