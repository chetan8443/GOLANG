//CHECKING THAT THE NUMBER IS POLINDROM OR NOT

package main

import "fmt"

func CheckPalindrom(num int) string {
	input_int := num
	res := 0
	var remainder int
	for num > 0 {
		remainder = num % 10
		res = (res * 10) + remainder
		num = num / 10
	}
	if input_int == res {
		return "number is polindrom"
	} else {
		return "number is not polyndrom"
	}

}
func main() {
	fmt.Println(CheckPalindrom(121))
	fmt.Println(CheckPalindrom(12321))
	fmt.Println(CheckPalindrom(1213))
}
