package main

import "fmt"
// created slice and map using struct

func main() {
	type s1 struct {
		colors string
	}
	type s2 struct {
		roll_no int
		name    string
	}
	
		//created slice by using struct 's2'

	s := []s2{{1, "arbind"}, {2, "yogesh"}}

	//created map by using struct 's2'
	m := map[int]s1{1: {"orange"}, 2: {"red"}, 3: {"yellow"}}
	
	for k,v:=range m{
	fmt.Printf("key : %d value : %s\n",k,v)
	}
	fmt.Println("")
	fmt.Println("below output is for Slice of struct")
	fmt.Println(s)
}
