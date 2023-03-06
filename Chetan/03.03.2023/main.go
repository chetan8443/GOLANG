package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"net/http"
	"net/url"
)

func main() {

	fmt.Println("=============WELCOME===============")
	// performGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func performGetRequest() {
	response, err := http.Get("http://localhost:8000/get") //asigning response to 'response' variable

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //for closing connection

	fmt.Println("Status code :", response.StatusCode) //printing status code
	fmt.Println("Content length :", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body) //reading responsebody
	// fmt.Println(string(content))                //converting response byte to string and print it

	var respString strings.Builder
	by, _ := respString.Write(content) //Write always returns len(content)
	fmt.Println(by)                    //printing length of content or byteCount

	fmt.Println(respString.String())
}

func PerformPostRequest() {
	var respString strings.Builder
	myurl := "http://localhost:8000/post"
	Requestbody := strings.NewReader(`
	 {
		"name":"CHETXY",
		"mobNo":30440343,
		"add":"bengalore"

	 }
	 `)

	resp, err := http.Post(myurl, "application/json", Requestbody)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	respString.Write(content)

	fmt.Println("Response after POST : ", respString.String())

}

func PerformPostFormRequest() {
	var respString strings.Builder
	myurl := "http://localhost:8000/postform"

	//Adding key-value pair for FORMDATA
	data := url.Values{}
	data.Add("NAME", "RAHUL")
	data.Add("PIN", "433212")
	data.Add("SURNAME", "TAHKRE")

	resp1, err := http.PostForm(myurl, data) //posting formdata & url
	if err != nil {
		panic(err)
	}

	defer resp1.Body.Close()
	content, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		panic(err)
	}
	respString.Write(content)
	fmt.Println("Content from POSTFORM :", respString.String())

}
