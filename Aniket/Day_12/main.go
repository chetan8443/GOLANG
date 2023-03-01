package main

import (
	"fmt"
)

func main() {

	Wg.Add(1)

	fmt.Println("Start of the Program.")
	GoRoutine("Without Go Routine")

	go GoRoutine("GO Routine")

	// time.Sleep(time.Second)

	fmt.Println("Program is still going")
	fmt.Println("Program Ends")

	fmt.Println(TestDefer())

	Wg.Wait()
}
