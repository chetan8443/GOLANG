
package main

import (
	"encoding/json"
	"fmt"
	"log"
)
// struct to store data in organised format
type student struct {
	Name string `json:"name"`
	Class string `json:"class"`
	RollNo int `json:"rollNo"`
}

func main()  {
	
	myvar:=student{
		Name: "abc",
		Class:"5th",
		RollNo: 23,
	}
// converting to json format
	jsonData,err:=json.Marshal(myvar)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Json Format:",string(jsonData))

	var newVar student
  // json to struct conversion
	err =json.Unmarshal(jsonData,&newVar)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unmarsheeld  Format:",newVar)
}
