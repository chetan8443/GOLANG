// Multiple Switch
package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter ID ")
	fmt.Scan(&a)
	switch a {
	case 100:
		var b string
		fmt.Print("Enter Password ")
		fmt.Scan(&b)
		switch b {
		case "password":
			fmt.Print("welcome user")
		default:
			fmt.Print("Invalid password")
		}
	default:
		fmt.Print("Invalid user")
	}
}

/* sample input and output

Input
Enter ID 100
Enter Password acdb

output
Invalid password*/
