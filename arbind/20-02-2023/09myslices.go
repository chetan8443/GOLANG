//creating a Slice and append elments into slice also sort the slice

package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Mango"}

	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[:2])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	//fmt.Println(highScores)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println(highScores)
}
