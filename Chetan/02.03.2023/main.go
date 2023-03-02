package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=afohaf"

func main() {
	//It parses a url into a URL structure.
	result, err := url.Parse(myurl)
	checkErr(err)

	//we can get specific value from result
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.Host)
	fmt.Println(result.RawQuery)

	//It parses RawQuery and returns the corresponding values
	fmt.Println("result of Query function in key,value")
	qparams := result.Query()
	fmt.Println(qparams)

	fmt.Println("For loop result")
	for _, v := range qparams {
		fmt.Println("Param is :", v)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
