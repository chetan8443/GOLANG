//write a Factorial Program using Recursion

package fact_reccursion

//function to calculate factorial using reccursion
func Factorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * Factorial(n-1) // factorial function calling itself
}
