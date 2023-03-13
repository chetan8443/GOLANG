package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	future := now.AddDate(-1, 2, 3)
	fmt.Println("One year and two days from now is:", future.Format("1995-02-03"))
}
