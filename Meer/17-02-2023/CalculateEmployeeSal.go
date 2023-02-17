package main

import "fmt"

func main() {

	var basicSal, hra, da, grossSal float64

	fmt.Print("Enter the Employee Basic Salary = ")
	fmt.Scanln(&basicSal)

	if basicSal <= 10000 {
		hra = (basicSal * 8) / 100
		da = (basicSal * 10) / 100
	} else if basicSal <= 20000 {
		hra = (basicSal * 16) / 100
		da = (basicSal * 20) / 100
	} else {
		hra = (basicSal * 24) / 100
		da = (basicSal * 30) / 100
	}

	grossSal = basicSal + hra + da
	fmt.Println("The Gross Salary of this Employee = ", grossSal)
}
