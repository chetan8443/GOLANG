// go routines
package main

import (
	"fmt"
	"time"
)

func main() {
	go greet("Hello")
	greet("Bhavana")
}

func greet(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}
