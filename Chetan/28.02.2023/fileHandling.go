package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("./myfile") //creating file
	if err != nil {
		panic(err) //if got an error execution of programme will stop
	}

	content := "content from myfile"
	length, err := io.WriteString(file, content) //writing content to file,output is 'length of file'

	//handling error
	if err != nil {
		panic(err)
	}
	fmt.Println("length :", length) //printing length of file

	defer file.Close()                        //Closing file
	fmt.Println(string(readFile("./myfile"))) //Calling readFile function & Converting output 'from dataByte to String'

}

func readFile(filename string) []byte {
	dataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return dataByte
}
