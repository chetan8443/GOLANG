// Taking input from the user
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our Pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

}
