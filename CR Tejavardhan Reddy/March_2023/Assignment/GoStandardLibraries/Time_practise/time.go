package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("Start")
	<-timer1.C
	fmt.Println("End")
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Dead Lock")
	}()
	time.Sleep(2 * time.Second)
	stop := timer2.Stop()
	if stop {
		fmt.Println("The end of the program")
	}
}

//func _()

// format the future date as a string
//formattedDate := futureDate.Format("02-Jan-2006")

// print the future date
