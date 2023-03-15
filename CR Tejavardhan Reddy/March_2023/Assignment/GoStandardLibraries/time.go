package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Print("\nThe Year is:\t", currentTime.Year())
	fmt.Print("\nThe Month is:\t", currentTime.Month())
	fmt.Print("\nThe Day is:\t", currentTime.Day())
	fmt.Print("\nThe Hour is:\t", currentTime.Hour())
	fmt.Print("\nThe Minute is:\t", currentTime.Minute())
	fmt.Print("\nThe Second is:\t", currentTime.Second())
	fmt.Print("\nThe time is:\t", currentTime)

}
