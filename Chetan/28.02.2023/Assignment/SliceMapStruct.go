package main

import "fmt"

func main() {
	type s1 struct {
		colors string
	}
	type s2 struct {
		roll_no int
		name    string
	}
	s := []s2{{1, "arbind"}, {2, "yogesh"}}

	m := map[int]s1{1: {"orange"}, 2: {"red"}, 3: {"yellow"}}
	
	for k,v:=range m{
	fmt.Printf("key : %d value : %s\n",k,v)
	}
	fmt.Println("")
	fmt.Println("below output is for Slice of struct")
	fmt.Println(s)
}
