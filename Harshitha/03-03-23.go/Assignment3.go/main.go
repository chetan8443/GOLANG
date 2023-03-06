package main

import "sync"

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go OddEven(&wg)

	go countAlpha(&wg)

	wg.Wait()

}