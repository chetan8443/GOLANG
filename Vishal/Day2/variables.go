package main

import "fmt"
  var input int
	var text string="Enter a number: "
	var b bool=false

func main() {
	
	_=b
	
	fmt.Print(text)
	fmt.Scanln(&input)
	output := reverse(input)
	fmt.Println("Reverse Number is :", output)

}

func reverse(num int) int {
	result := 0
	for num != 0 {
		var rem int = num % 10
		result = result*10 + rem
		num = num / 10
	}

	return result
}
