package main

const testURL = "https://google.com"
const myurl = "https://lco.dev:3000/tutcss?Coursename=python&id=234096"

func main() {

	// Calling below function to get response body from the given url
	getResponse(testURL)

	// Calling below function to get various parameter out og response body
	urlResp(myurl)

	// Calling below function to get the Query parameter from given url
	getQueryParam(myurl)
}
