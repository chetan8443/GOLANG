package main

import (
	"fmt"
	"os"
)

// Function to check working of "defer" keyword
func deferExa() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Function to write the contents of the string variable to file
func writeFile(content string, file string) {
	err := os.WriteFile(file, []byte(content), 0666)
	checkNilError(err)
}

// Function to read the contents of the given file
func readFile(file string) string {
	bs, err := os.ReadFile(file)
	checkNilError(err)
	return string(bs)
}

// Function to check whether error is present or not
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
