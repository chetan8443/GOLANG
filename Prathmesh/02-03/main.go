package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//go greeter("hello")
	//greeter("world")

	webitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range webitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

//	func greeter(s string) {
//		for i := 0; i < 6; i++ {
//			time.Sleep(3 * time.Millisecond)
//			fmt.Println(s)
//		}
//	}
func getStatusCode(endpoint string) {

	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("oops in endpoint")

	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
