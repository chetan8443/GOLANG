//Program to check if a particular key value pair is present in a map
package main

import "fmt"

func main() {
	var key int
	var found byte
	fmt.Println("Enter key to be checked")
	fmt.Scan(&key)

	//declaration and initialization of the map with key value pairs
	map1:= map[int]string{
        10: "Rohit",
        11: "Kohli",
        12: "Ashwin",
        13: "Sundar",
        14: "Samson",
		15: "Jadeja",
    }
	//iterating along the map to check for the key
	for k , v := range map1 {
		if k == key{
			fmt.Println("Key",key,"with value",v,"found ")
			found = 'f'
			break
		}
	}
	if found !='f'{
		fmt.Println("Key ", key,"not found")
	}
}