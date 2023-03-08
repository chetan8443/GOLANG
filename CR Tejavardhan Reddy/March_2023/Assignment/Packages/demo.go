package main

import (
	"encoding/json"
	"fmt"
)

var teja struct {
	inte int
}

func main() {
	var slice = make([]int, 0)
	fmt.Println("Golang packages")
	str_marshal, _ := json.Marshal("String")
	float_marshal, _ := json.Marshal(3.14)
	struct_marshal, _ := json.Marshal(teja)
	fmt.Println("After marshalling", str_marshal, float_marshal, struct_marshal, slice)
	// var struct_unmarshal teja
	// ba:=[]byte(teja.inte)
	// json.Unmarshal(ba, &struct_unmarshal)
	json.Unmarshal()
}
