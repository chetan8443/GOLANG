package main
import (
	"fmt"
	"time"
)
func main(){
	actuvalDate := time.Now()  
	dateafteryear := actuvalDate.AddDate(1, 0, 2) 
	fmt.Println("The date after adding one year and two days from today is:", dateafteryear.Format("2006-02-01")) 
}
