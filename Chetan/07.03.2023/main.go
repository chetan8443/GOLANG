//CREATING JSON DATA
package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"CourseName"`
	Price    int      `json:"price"`
	Password string   `json:"-"`				//'-' for Hide values for security purpose
	Tag      []string `json:"tag,omitempty"`	//omitempty to ignore error if value is empty
}

func main() {
	
// 	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	CdacCourse := []Course{
		{Name: "Python", Price: 192021, Password: "saveit", Tag: []string{"web-dev", "js"}},
		{Name: "Java", Price: 192021, Password: "saveit", Tag: nil},
	}
	// dataByte, err := json.Marshal(CdacCourse)		//Marshal returns the JSON encoding of 'CdacCourse'.
	dataByte, err := json.MarshalIndent(CdacCourse, "", "\t") //applies Indent to format the output. Each JSON element in the output will begin on a new line beginning with prefix followed by one

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", dataByte) //It's printing readable format Json Data
	// fmt.Println(dataByte)    It's printing ASCII values

}

func DecodeJson() {
	var course Course
	jsonData := []byte(`
		{
		"courseName":"PythonDeveloper",
		"price":299,
		"Password":"4324",
		"Tag":["web-dev","js"]}
	`)
	checkValid := json.Valid(jsonData)

	if checkValid {
		json.Unmarshal(jsonData, &course)
		fmt.Printf("%#v",course)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
}
