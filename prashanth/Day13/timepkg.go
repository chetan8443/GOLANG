package main 

import (
	 "fmt"
     "os"
	 "time"
)	

func Create() {
	a, err:= os.Create("test.txt") // iam create a empty file
	if err !=nil{
		fmt.Println(err)
	}
	b, err:= a.WriteString("Hi Everyone") // here i am creating a string it will open txt file
	fmt.Println(b)
}
func  main()  {
	t:= time.Now()   // here i am mention the time
	fmt.Println(t)

	s:= time.Date(1991, time.August, 1, 1, 1, 1, 1, time.UTC) // here iam creating time and Date
	fmt.Println(s)
	
}
