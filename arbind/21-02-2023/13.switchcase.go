package main

import ( //Importing libraries
	"fmt"
	"math/rand"
	"time"
)

func main() { // creating a mainn function
	fmt.Println("Switch and case in ") //printing

	rand.Seed(time.Now().UnixNano()) //Generating Random number
	diceNumber := rand.Intn(6) + 1   // random number will be generating in 0 to 6

	fmt.Println("Vlaue of dice is : ", diceNumber)

	switch diceNumber { //swich statement
	case 1: // case according to requirements
		fmt.Println("Dice value is 1 so you can open")
	case 2:
		fmt.Println("Dice value is 2 so you can't open")
	case 3:
		fmt.Println("Dice value is 3 so you can't open")
	case 4:
		fmt.Println("Dice value is 4 so you can't open")
	case 5:
		fmt.Println("Dice value is 5 so you can't open")
	case 6:
		fmt.Println("Dice value is 6 so you can open")
	}
}
