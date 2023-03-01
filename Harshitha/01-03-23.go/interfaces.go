package main

import "fmt"

//no fields
type englishbot struct{}
type spanishbot struct{}

func main(){
	
}
func (eb englishbot) getGreeting() string{
	//Very custo logic for english logic
	return "Hi David!"
}

func (sb spanishbot) getGreeting() string{
	return "Hallo!"
}