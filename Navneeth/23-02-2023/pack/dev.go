package pack

var g = 555 //global variable

//function to reverse a string
func Rev(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

//function to print global variable
func Global() int {
	return g
}
