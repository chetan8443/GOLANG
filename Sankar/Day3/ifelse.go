package main 
import "fmt"
func main() {
	
	for i:=0;i<=10;i+=1 {
		if i == 0{
			fmt.Println("zero value: ", i)
		} else if i<=3 && i>=1{
			fmt.Println("value between (1-3)", i)
		} else if i<=8 {
			fmt.Println("greater than 3:", i)
		}else {
			break
		}
		}
	}
