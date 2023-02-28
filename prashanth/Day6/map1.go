package main

import "fmt"
func main()  {
	//syntax1--------->
	 players:=map[string]string{
	
		"player1": "Sachin",
		"player2": "Dhoni",
	}
	//syntax2--------->
	 //var players map[string]string := map[string]string{

        //"player1": "Sachin",
		//"player2": "Dhoni",
	//}
	fmt.Println(players)
	fmt.Println(players["player1"])
	fmt.Println(players["player2"])
}
