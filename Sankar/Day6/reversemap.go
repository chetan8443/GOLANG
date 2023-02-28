package main

import "fmt"
func main(){
		
	reversemap()

}

func reversemap(){
	emplist := map[int]string {1001:"sam",1002:"maharaj",1003:"jhon"}
	fmt.Println(emplist,"\n")
	list1 := []int {}
	list2 := []string {}
	for i,j := range emplist{
		list1 = append(list1,i)
		list2 = append(list2,j)
	}
	newemplist := map[string] int {}
	for i:=0;i<len(emplist);i++{
		newemplist[list2[i]] = list1[i]
	}
	fmt.Println(newemplist,"\n")
}
