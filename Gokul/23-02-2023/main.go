// Pgm to test importing of multiple packages. Main.go will call test.go inside hope folder
package main

import (
	"fmt"
	a "funct/hope" //funct is the module inside go.mod file
)

func main() {
	fmt.Println("I'm main")
	//Calling test.go inside hope
	a.Test()
	//Accessing global variable inside function glb()
	fmt.Println("Global var is", glb())
}
