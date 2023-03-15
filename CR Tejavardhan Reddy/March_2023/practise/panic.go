package main

import "os"

func main() {
	panic("Error Situation")
	_, err := os.Open("./file.go")
	if err != nil {
		panic(err)
	}
}
