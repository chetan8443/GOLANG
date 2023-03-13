package main

import (
	"fmt"
	"time"
)

func main() {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	t := time.Now()
	fmt.Println("Current time in New York:", t.In(loc).Format(time.RFC3339))

	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	fmt.Println("Current time in Tokyo:", t.In(loc).Format(time.RFC3339))
}
