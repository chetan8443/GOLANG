// program to convert 12 hour format time into 24 hour format
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "07:05:45PM"
	s1 := "12:01:00PM"
	s2 := "12:01:00AM"
	s3 := timeConversion(s)
	s4 := timeConversion(s1)
	s5 := timeConversion(s2)

	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}

func timeConversion(s string) string {
	// Write your code here
	s = strings.ToLower(s)
	pm := strings.HasSuffix(s, "pm")
	am := strings.HasSuffix(s, "am")

	t := s[:len(s)-2]

	timeArr := strings.Split(t, ":")
	hh, mm, ss := timeArr[0], timeArr[1], timeArr[2]
	hhInt, err := strconv.Atoi(hh) // string to int

	if err != nil {
		panic(err.Error())
	}

	if am && hhInt == 12 {
		hhInt = 0
	}

	if pm && hhInt != 12 {
		hhInt += 12
	}

	hh = strconv.Itoa(hhInt) // int to string

	// output: 07:15:45 or 23:03:33
	return fmt.Sprintf("%02s:%02s:%02s", hh, mm, ss)

}
