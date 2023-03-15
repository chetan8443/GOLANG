package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// write()
	read()
}
func write() {
	file, err := os.Create("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hi..there")
	file.WriteString("\nWelcome to the Coding of go")
	file.Close()
}
func read() {
	data, err := ioutil.ReadFile("demo.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	readstring := string(data)
	fmt.Println(readstring)
}
