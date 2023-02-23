package main 
import "fmt"
func main() {
	department := "str"
	switch department {
	case "fin":
		fmt.Println("Finance department")
	case "str":
		fmt.Println("Storage department")
	case "opt":
		fmt.Println("Operations department")
	default:
		fmt.Println("Please enter VALID DEPARTMENT Name")
	}
    
}
