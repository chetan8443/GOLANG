package main

import (
	"fmt"
	"time"
)

func main() {
	var dobstr string
	fmt.Println("Enter the DOB ")
	fmt.Scan(&dobstr)

	dob, err := time.Parse("01-02-2006", dobstr)
	if err != nil {
		panic(err)
	}

	currDate := time.Now()
	//currDateStr := currDate.Format("01-02-2006")

	year := currDate.Year()
	dobYear := dob.Year()

	month := currDate.Month()
	dobMonth := dob.Month()

	day := currDate.Day()
	dobDay := dob.Day()

	yearDiff := year - dobYear

	var monthDiff int
	if month > dobMonth {
		monthDiff = int(month - dobMonth)
	} else {
		monthDiff = int(dobMonth - month)
	}
	dayDiff := day - dobDay

	fmt.Println("The Age is : ", yearDiff, " years ", int(monthDiff), " Months ", dayDiff, " Days Year Old")

}
