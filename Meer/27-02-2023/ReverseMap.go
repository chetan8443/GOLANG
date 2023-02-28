package main

import "fmt"

func main() {

	marks := map[string]int{
		"Meer":  95,
		"Fahad": 86,
		"Ali":   90,
		"Syed":  75,
	}

	fmt.Println("\nThe elements in map before : ")
	list := []string{}
	for k, v := range marks {
		list = append(list, k)
		fmt.Println(k, v)

	}

	fmt.Println("\nThe elements in map after reverse : ")
	for i := len(list) - 1; i >= 0; i-- {
		fmt.Println(list[i], marks[list[i]])
	}
}
