package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	user := User{
		Name: "Aniket",
		Age:  27,
		Address: Address{
			City:    "Pune",
			State:   "Maharashtra",
			Country: "India",
		},
	}

	bs, err := json.MarshalIndent(user, " ", "   ")
	CheckNullValue(err)
	jsonData := string(bs)
	fmt.Println(jsonData)

	var user1 User
	err = json.Unmarshal([]byte(jsonData), &user1)
	CheckNullValue(err)

	fmt.Println(user1)

	var user2 User
	jsonBs, err := ioutil.ReadFile("data.json")
	CheckNullValue(err)
	err = json.Unmarshal(jsonBs, &user2)
	CheckNullValue(err)

	fmt.Println(user2)
}
