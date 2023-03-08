package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
	LoggedAt time.Time
}

func main() {

	user := User{}
	user.Id = 1122
	user.Name = "Nav"
	user.Password = "golang"
	user.LoggedAt = time.Now()
	//Writing struct type to a JSON file
	content, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("userfile.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//Reading into struct type from a JSON file
	content, err = ioutil.ReadFile("userfile.json")
	if err != nil {
		log.Fatal(err)
	}
	user2 := User{}
	err = json.Unmarshal(content, &user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Id:%d, Name:%s, Password:%s, LoggedAt:%v \n", user2.Id, user2.Name, user2.Password, user2.LoggedAt)
}
