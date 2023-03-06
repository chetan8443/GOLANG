// The cats are at positions 2 (Cat A) and 5 (Cat B), and the mouse is at position 4. Cat B,
// at position 5 will arrive first since it is only 1 unit away while the other is 2 units away. Return 'Cat B'.

// If cat A catches the mouse first, print Cat A.
// If cat B catches the mouse first, print Cat B.
// If both cats reach the mouse at the same time, print Mouse C as the two cats fight and mouse escapes.

package main

import "fmt"

func main() {
	fmt.Print("Enter the values:")
	var x, y, z int
	fmt.Scanln(&x, &y, &z)
	fmt.Println(logic(x, y, z))

}
func logic(x, y, z int) string {
	x1 := x - z
	x2 := y - z
	if x1 < 0 {
		x1 *= -1
	}
	if x2 < 0 {
		x2 *= -1
	}
	if x1 == x2 {
		return "Mouse C"
	} else if x1 < x2 {
		return "Cat A"
	} else {
		return "Cat B"
	}
}
