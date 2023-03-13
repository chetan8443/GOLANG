package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	mydata := []byte("Meer is writing some text in the file & also fetching it using golang code !!!")

	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("myfile.txt", mydata, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("myfile.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

}
