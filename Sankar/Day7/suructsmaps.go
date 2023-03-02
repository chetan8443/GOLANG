package main
import "fmt"
type Sports struct {
	Name string
	Tool string
}

func main() {
	s1:= Sports{"Cricket","Bat"}
	s2:= Sports{"Tennis","Rocket"}

	Count := map[int]Sports{
		1: s1,
		2: s2,
	}
	for i:=1;i<=len(Count);i++{
		fmt.Println(Count[i].Name,Count[i].Tool)
	}
	}
	
	
