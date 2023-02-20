package second

import "fmt"

func Pattern2() {
	rows := 5

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k != (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
