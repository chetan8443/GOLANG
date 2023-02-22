package main // Package name

import "fmt" // module name

func main() {
	threshold := 60 //initialized the value for variable threshold here
	if threshold < 50 {
		fmt.Println("Do not turn of the switch") // if else conditions.
	} else if threshold == 50 {
		fmt.Println("Notify the operator")
	} else if threshold > 90 {
		fmt.Println("turn of the switch")
	} else {
		fmt.Println("ignore")
	}

}
