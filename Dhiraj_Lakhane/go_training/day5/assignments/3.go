//Write a program that declares a slice of strings and sorts it in alphabetical order.
package main
 
import (
	"fmt"
	"sort"
)

func sortString(){
	str:=[]string{"dhiraj","arbind","vishal","nayan"}

	fmt.Print("before sort: ")
	fmt.Println(str)
	fmt.Print("after sort: ")
	sort.Strings(str)
	fmt.Println(str)



}