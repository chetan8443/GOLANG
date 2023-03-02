package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	response, err := http.Get(url)

	CheckError(err)
	defer response.Body.Close() //it closes connection

	databytes, err := ioutil.ReadAll(response.Body)

	CheckError(err)

	content := string(databytes)
	fmt.Println(content)

}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
