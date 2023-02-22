//Pgm to check if entered character is a vowel,using multiple switch case statements.
package main

import "fmt"

//Main function
func main() {
	var ch string
	fmt.Println("Enter a character (alphabet)")
	fmt.Scan(&ch)

	//Switch case 
	switch ch {
		case "a","e","i","o","u","A","E","I","O","U":
			fmt.Println(ch," is a vowel")
		default :
			fmt.Println(ch,"is a consonant")
	}
}