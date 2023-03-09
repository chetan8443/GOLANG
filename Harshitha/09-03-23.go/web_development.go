package main

import(
	"fmt"
	"net/http"
)
func uploadFile(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Uploading file Harshitha")
}

func setupRoutes(){
	http.HandleFunc("/upload",uploadFile)
	http.ListenAndServe(":8080",nil)
}

func main(){
	fmt.Println("Go file upload tutorial")
	setupRoutes()
}