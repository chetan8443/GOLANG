package main 
import (
	"encoding/json"
	"fmt"
)
type book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json: "author"`
	Quantity int `json:"quantity"`
}
func main(){
	books := book{ID: "1",Title: "In Search of Lost Time",Author: "Marcel Proust", Quantity:2}
	jsonbooks,err := json.Marshal(books)
	if err !=nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(jsonbooks))
	var books_new book
	err = json.Unmarshal(jsonbooks,&books_new)
	if err != nil {
		fmt.Println("error ! ",err)
	}
	fmt.Println(books_new)
}
