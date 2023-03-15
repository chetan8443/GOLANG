package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("1st line in main")
	wg.Add(2)

	go func2()
	go func1()
	wg.Wait()
	fmt.Println("Last line in main")
}

func func1() {
	fmt.Println("Function 1")
	wg.Done()
}
func func2() {
	fmt.Println("Function 2")

	wg.Done()
}

// func func3() {
// 	fmt.Println("Function 3")
// }
