package constant_var

import "fmt"

const str string = "Infobell" // constant string type varaiable
const pi float64 = 3.14       // constant float type varaiable

func Const_var() {
	const b bool = false // constant boolean type varaiable
	const i int = 10     // constant integer type varaiable

	fmt.Println("constant string :", str) // printing global constant string type variable
	fmt.Println("constant float :", pi)   // printing global constant float type variable
	fmt.Println("constant bool :", b)     // printing local constant bool type variable
	fmt.Println("constant integer :", i)  // printing local constant integer type variable

}
