package main

import (
	"fmt"
)

func main() {
	fmt.Println("type 'User' or 'Admin' : ")
	var role string
	fmt.Scan(&role)


	switch role {
		
			case "Admin":
				fmt.Println("You loggin as a ADMIN")
				fmt.Println("Enter a number between 11 or 22 or 33 ")
				var num int
				fmt.Scan(&num)

							switch num {
										case 11:
											fmt.Println("Value is 11")

										case 22:
											fmt.Println("Value is 22")

										case 33:
											fmt.Println("Value is 33")

										default:
											fmt.Println("Out of range")

										}
									case "User":
										fmt.Println("You loggin as a USER")
										fmt.Println("Enter a number between 1-3 ")
										var num int
										fmt.Scan(&num)
										switch num {
										case 1:
											fmt.Println("Value is 1")

										case 2:
											fmt.Println("Value is 2")

										case 3:
											fmt.Println("Value is 3")
										default:
											fmt.Println("Out of range")

							}

	}

}
