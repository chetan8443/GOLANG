package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.om",
	}
	c:=make(chan string)
  
  //for loop  for CheckLinks function calls as much as we have links loop will call CheckLink function 
	for _, link := range links {
		go checkLink(link,c)
	}
  
  //for loop for printing channel message
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

 
 
}

func checkLink(link string,c chan string) {
	_, err := http.Get(link)
	if err!=nil{

		fmt.Println(link,"is might down !")
		c<-"yep it's  might down!"
		return
	}
	fmt.Println(link,"is Working !")
	c<-"yep it's up !"
}
