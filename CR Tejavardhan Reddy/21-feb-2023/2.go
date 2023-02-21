/*
Go lang code to determine the even places and odd places in a string seperatedly
Input:tejavardhan
Output:
tjvrhn
eaada
*/

package main

import "fmt"

func main1() {
	fmt.Print("Enter the string : ")
	var S string
	fmt.Scanln(&S)
	oddCharacter(S)
	evenCharacters(S)
}

func oddCharacter(S string) {
	//var res1 string
	fmt.Print("The odd places are : ")
	for i := 0; i < len(S); i += 2 {
		fmt.Printf("%c", S[i])
	}
	fmt.Print("\n")

}
func evenCharacters(S string) {
	fmt.Print("The even places are : ")
	for i := 1; i < len(S); i += 2 {
		fmt.Printf("%c", S[i])
	}
}
