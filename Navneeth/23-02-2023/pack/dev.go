package pack

var g = 555

func Rev(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func Global() int {
	return g
}
