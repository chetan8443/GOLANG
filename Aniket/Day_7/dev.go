package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var signal []string
var list []int

// Function to get Statuscode from the list of URLs
func getStatusCode(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Something wrong in website")
	} else {
		mutex.Lock()
		signal = append(signal, url)
		mutex.Unlock()
		fmt.Printf("%d Status OK code for %s\n", res.StatusCode, url)
	}
}

// Function to check GoRoutine and Mutex function
func routine1(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	fmt.Println("1st Go Routine")
	mutex.Lock()
	list = append(list, 1)
	mutex.Unlock()
}

// Function to check GoRoutine and Mutex function
func routine2(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	fmt.Println("2nd Go Routine")
	mutex.Lock()
	list = append(list, 2)
	mutex.Unlock()
}

// Function to check GoRoutine and Mutex function
func routine3(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	fmt.Println("3rd Go Routine")
	mutex.Lock()
	list = append(list, 3)
	mutex.Unlock()
}
