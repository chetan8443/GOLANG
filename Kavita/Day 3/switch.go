package main //package name

import "fmt" // module name

func main() { // function started here
	letters := "C"
	switch letters { // switch case begins from here
	case "A":
		fmt.Println("Print this for A")
	case "B":
		fmt.Println("Print this for B")
	case "C", "D":
		fmt.Println("Print this for C and D")
	default:
		fmt.Println("Print this otherwise")
	}
}
