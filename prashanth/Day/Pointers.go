package main
import "fmt"
func main() {
	var p int = 45
	pointer_p:=&p
	fmt.Println(p)
	fmt.Println(pointer_p)
	fmt.Println(*pointer_p)
	*pointer_p = 900
	fmt.Println(*pointer_p)
	p = 5252
	fmt.Println(*pointer_p)
}
