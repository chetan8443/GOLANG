package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("The Local Time:", currentTime)
	fmt.Println("The Standard Time:", currentTime.UTC())
}
