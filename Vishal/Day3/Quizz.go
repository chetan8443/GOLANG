package main

import "fmt"
var score int =0 //global Variable to use this in all program
func main() {
	start() // function will do all task
	if(score<=2){
		fmt.Print("You need to take quizzz again")
	}else{
		if(score==5){
			fmt.Println("Excellent you scored 5 out of 5")
		}else
		{
			fmt.Printf("\n Excellent you  scored %d out of 5",score)
			fmt.Println("You can do better take quizz again")
		}
	}
}

func start() {
	fmt.Println("Q1 Does Go support type inheritance?")
	fmt.Println("Select one Option")
	fmt.Println("1) Yes")
	fmt.Println("2) No")
	var input int
	fmt.Scanln(&input)    // to take input and send to switch 
	switch input {
	case 1:fmt.Println("opps! Wrong answer")
	case 2:fmt.Println("Good Job correct answer ! ")
	         score++

		
	}

	fmt.Println("Q 2 - Which of the following is a derived type in Go?")
	fmt.Println("Select one Option")
	fmt.Println("1)  Union types")
	fmt.Println("2) Function types")
	fmt.Println("3) Slice types")
	fmt.Println("4) All of the above.")
	var input2 int
	fmt.Scanln(&input2)  // to take input and send to switch 
	switch input2 {
	case 1:fmt.Println("opps! Wrong answer")
	case 2:fmt.Println("opps! Wrong answer")
	case 3:fmt.Println("opps! Wrong answer")
	case 4:fmt.Println("Good Job correct answer ! ")
	score++

		
	}
	fmt.Println("Q3 Which of the following transfers control to the labeled statement in Go?")
	fmt.Println("Select one Option")
	fmt.Println("1)  break")
	fmt.Println("2) continiue")
	fmt.Println("3) go to")
	fmt.Println("4) None of the above.")
	var input3 int
	fmt.Scanln(&input3)  // to take input and send to switch 
	switch input3 {
	case 1:fmt.Println("opps! Wrong answer")
	case 2:fmt.Println("opps! Wrong answer")
	case 3:fmt.Println("Good Job correct answer ! ")
	score++
	case 4:fmt.Println("opps! Wrong answer")

		
	}

	fmt.Println("Q4 In Go language, variables of different types can be declared in one statement.")
	fmt.Println("Select one Option")
	fmt.Println("1) True")
	fmt.Println("2) False")
	var input4 int
	fmt.Scanln(&input4)  // to take input and send to switch 
	switch input4 {
	case 1:fmt.Println("Good Job correct answer ! ")
	score++
	case 2:fmt.Println("opps! Wrong answer")

		
	}
	fmt.Println("Q4 Can you define a pointer to pointer in Go?")
	fmt.Println("Select one Option")
	fmt.Println("1) True")
	fmt.Println("2) False")
	var input5 int
	fmt.Scanln(&input5)
	switch input5 {
	case 1:fmt.Println("Good Job correct answer ! ")
	score++
	case 2:fmt.Println("opps! Wrong answer")

		
	}

}
