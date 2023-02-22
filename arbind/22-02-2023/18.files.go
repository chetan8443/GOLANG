// Write a content into file and read it from the file

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() { // creating main function
	fmt.Println("Welcome to files in golang") // Printing a statement

	content := "This need to go in a file -LearnCodeOnline.in" // content that we want write into a file

	file, err := os.Create("./mylcogofile.txt") // passing location and filename

	checkNilErr(err) // checking error

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is : ", length)
	defer file.Close()            // using defer to close the file at the end
	readFile("./mylcogofile.txt") // calling readFile function inside main function

}

func readFile(filename string) { // creating readFile function ouside the main function
	databyte, err := ioutil.ReadFile(filename) //reading the content in byte format
	checkNilErr(err)

	fmt.Println("Text data is \n", string(databyte)) //Printing the databye
}

func checkNilErr(err error) { //creating a function for checking error
	if err != nil {
		panic(err)
	}

}
