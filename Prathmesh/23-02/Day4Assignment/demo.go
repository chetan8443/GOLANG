// Find Out How do you call a global variable in other golang files.
package main

import "fmt"

var Second string = "global" //initializing the gloabal variable

func demo() { //demo function for calling in other file
	fmt.Println("calling demo function in main file")
}
