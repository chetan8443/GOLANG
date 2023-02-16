package main

import "fmt"

func main() {
	fmt.Print("Let check number is armstrong or not : ")
	var num int
	var sum int = 0

	fmt.Scan(&num)
	temp:=num
	for num > 0 {
		r := num % 10
		d:=sum+(r * r * r)
		sum=d
		num = num / 10
		

	}

	if sum == temp {
		fmt.Print("armstrong  number")
	} else {
		fmt.Print("not armstrong  number")
	}

}
