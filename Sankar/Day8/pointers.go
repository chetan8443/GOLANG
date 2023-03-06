package main
import "fmt"
func main() {
	var a int = 14
	pointer_a:=&a
	fmt.Println(a)
	fmt.Println(pointer_a)
	fmt.Println(*pointer_a)
	*pointer_a = 765
	fmt.Println(*pointer_a)
	a = 2344
	fmt.Println(*pointer_a)
}

