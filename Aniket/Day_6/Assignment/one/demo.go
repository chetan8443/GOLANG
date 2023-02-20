package one

func FactR(num int) int {
	fact := 1
	if num == 1 {
		return 1
	} else {
		fact = num * FactR(num-1)
	}
	return fact
}
