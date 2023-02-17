package main

import "fmt"

func main() {
	notes := [8]int{500, 100, 50, 20, 10, 5, 2, 1}
	var amount int
	fmt.Print("Enter the Total Amount of Cash = ")
	fmt.Scanln(&amount)

	temp := amount
	for i := 0; i < len(notes) ; i++ {
		fmt.Println(notes[i], " Notes = ", temp/notes[i])
		temp = temp % notes[i]
	}
}
