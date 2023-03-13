package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	//finding out date one year and two days later
	oneYearLater := t.AddDate(1, 0, 2)
	//formating time to DD-MM-YYYY
	now := t.Format("02-01-2006")
	year := oneYearLater.Format("02-01-2006")
	//printing date
	fmt.Printf("Time Now: %v\n Time one year two days later: %v\n", now, year)

	//time zone for America/New York
	loc, _ := time.LoadLocation("America/New_York")
	zone := time.Now().In(loc)
	//time zone for Europe/London
	location, _ := time.LoadLocation("Europe/London")
	AnotherZone := time.Now().In(location)
	//printing time in two different time zones
	fmt.Printf("Time in America/New York : %v \n Time in Europe/London : %v \n", zone, AnotherZone)

}
