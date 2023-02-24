package main

import "fmt"

func main() {
	fmt.Println("code for nested ifelse and using break and continue")
	var n int
	fmt.Println("Enter the marks you have got u will get your grade")
	fmt.Scanln(&n)

	if n <= 35 {
		fmt.Println("You are Failed")
		if n >= 30 {
			fmt.Println("you are eligible for re-examination")

		}

	} else if n > 35 && n < 60 {
		fmt.Println("you got C grade")
	} else if n >= 60 && n < 69 {
		fmt.Println("you got B grade")
	} else if n >= 70 && n < 84 {
		fmt.Println("you got A grade")
	} else if n >= 85 && n < 100 {
		fmt.Println("you got A++ grade")
	} else {
		fmt.Println("Please enter valid Number to getting output")

	}

}
