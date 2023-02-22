//Pgm to print nos from 1 to num(entered by user)
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the num")
	fmt.Scan(&num)
	fmt.Println("Numbers from 1 to",num,"are ")
	//for loop
	for i:=1;i<=num;i++{
		fmt.Println(i)
	}
}