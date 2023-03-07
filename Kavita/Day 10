package main

import ("fmt"
		"encoding/json"
		"log"
		"net/http"
)


type employee struct{
		name string
		ID int
	}

var emp = []employee{
		{name: "kavita", ID: 100}, 
		{name: "Alka", ID: 101},
	}


func main(){
	http.HandleFunc("/emp/", getempbyID)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getempbyID	(w http.ResponseWriter, r *http.Request){

	id := r.URL.Path[len("/emp/"):]
	for _, emp := range emp {

		if fmt.Sprintf("%d", emp.ID) == id {
		json.NewEncoder(w).Encode(emp)
		return
		}
	}
	http.NotFound(w,r)

}
