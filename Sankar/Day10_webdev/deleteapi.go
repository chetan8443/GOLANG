package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
    "github.com/gorilla/mux"
)
type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage! ")
    fmt.Println("Endpoint Hit: homePage")
}
func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
    r := mux.NewRouter()

    http.HandleFunc("/", homePage)
	http.HandleFunc("/Articles", returnAllArticles)
    r.HandleFunc("/Articles/{Title}", deleteEmployee).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8768", nil))
}
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    Title := vars["Title"]
    for index, formula := range Articles {
        if formula.Title == Title {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }
}

var Articles []Article
func main() {
    Articles = []Article{
        Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}
