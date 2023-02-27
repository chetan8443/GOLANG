package main

import ("fmt"
"strings")

func search(element) bool{
	for i=0;i<len(counter);i++{
		if counter[i]==element{
			return true
		}else{
			return false
		}
	}

}

func main(){
	var a string="apple banana"
	var b string="banana"
	var element string
	h:=strings.Split(a," ")
	t:=strings.Split(b," ")
	var i int=0
	var counter = map[string] int {}
	for i=0;i<len(h);i++{
		if search(h[i]){
			counter[i]=counter[i]+1
		}
	}
	for i=0;i<len(t);i++{
		if search(t[i]){
			counter[i]=counter[i]+1
		}
	}

}