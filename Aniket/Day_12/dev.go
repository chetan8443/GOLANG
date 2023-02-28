package main

import (
	"fmt"
	"sync"
)

var Wg = new(sync.WaitGroup)

func GoRoutine(str string) {
	defer Wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Println(str)
	}
}

func TestDefer() string {
	defer fmt.Println("Defer Executes")
	return "Function Ends"
}

// go func(msg string) {fmt.Println(msg)}("Going")
