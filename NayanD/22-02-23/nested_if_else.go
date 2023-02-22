package main

import "fmt"

func main() {

	var balance int

	fmt.Println("Enter your balance: ")
	fmt.Scanf("%d", &balance)

	if balance <= 0 {
		fmt.Printf("Your balance is low ,Please add some funds or penalty will be charged ")

	} else if balance >= 0 && balance <= 500 {
		fmt.Printf("your balance is %v , add funds minimum balance 500 is required", balance)

	} else {

		fmt.Printf("your balance is %v \n Thankyou!!!", balance)
	}

}
