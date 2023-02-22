package main
import "fmt"
var commit_msg = "variable declaration"
var commit_num = 2


func main() {
	fmt.Println("the description for the commit : ", commit_msg)
	fmt.Printf("this is the %vnd commit!\n",commit_num)
	fmt.Printf("the datatype of commit_msg is := %T : and the datatype fo commit_num is := %T : \n",commit_msg,commit_num)
	commit_num := 3
	commit_error:=" "
	fmt.Println("changed commit number : ",commit_num)
	fmt.Println("the error follow as :",commit_error )
}
