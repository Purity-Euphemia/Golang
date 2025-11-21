package pasince

func Swap(a *int, b *int){
	value := *a
	*a = *b
	*b = value
}