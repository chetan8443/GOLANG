package first

import "fmt"

func Pattern() {

	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}

}
