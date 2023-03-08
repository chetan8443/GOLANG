package main

import "fmt"

func main() {
  mymap := map[string]int{
    "Elliot": 25,
    "Sophie": 24, 
    "Fraser": 20,
  }
  //There is no indexing in map
  for i :=range(mymap){
	for j:=range(mymap){
		if mymap[i]>mymap[j]{
			mymap[i]=mymap[j]
		}
	}
  }
  fmt.Println(mymap)
}