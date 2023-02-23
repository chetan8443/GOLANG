package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Function to get response from the given url
func getResponse(url string) {
	res, err := http.Get(url)
	checkNilError(err)
	bs, err := ioutil.ReadAll(res.Body)
	checkNilError(err)
	fmt.Println(string(bs))
	defer res.Body.Close()
}

// Function to get the various parameter out of url response body
func urlResp(myurl string) {
	res, err := url.Parse(myurl)
	checkNilError(err)
	fmt.Println("The Host Of URL is", res.Host)
	fmt.Println("The Path of URL is", res.Path)
	fmt.Println("The Scheme of URL is", res.Scheme)
	fmt.Println("The RawQuery of URL is", res.RawQuery)
	fmt.Println("The Port of URL is", res.Port())
}

// FUnction to get the Query response parameter from the given url
func getQueryParam(myurl string) {
	res, err := url.Parse(myurl)
	checkNilError(err)
	param := []string{"Coursename", "id"}
	values := res.Query()
	for _, i := range param {
		fmt.Printf("The parameter %s has value %s\n", i, values.Get(i))
	}
}

// Function to check whether the required function is sending error as response or not
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
