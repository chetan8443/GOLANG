// //In the year  = 2017, January has 31 days, February has 28 days,
// March has 31 days, April has 30 days, May has 31 days, June has 30 days, July has 31 days,
// and August has 31 days. When we sum the total number of days in the first eight months,
// we get 31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 = 243. Day of the Programmer is the 256th day,
//
//	so then calculate 256 - 243 = 13 to determine that it falls on day 13 of the 9th month (September).
//	We then print the full date in the specified format, which is 13.09.2017.
//
// Sample Input 1
// 2016
// Sample Output 1
// 12.09.2016
package main

import "fmt"

func main() {
	var year int
	var totalDays int
	fmt.Println("Enter the year:")
	fmt.Scanln(&year)
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		totalDays = 244
	} else {
		totalDays = 243
	}
	day := 256 - totalDays
	fmt.Print(day, ".09.", year)
}
