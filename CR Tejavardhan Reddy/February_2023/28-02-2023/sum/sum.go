package teja

func Sum(arr []int) int {
	sum := 0
	for _, element := range arr {
		sum = sum + element
	}
	return sum
}
func EvenSum(arr []int) int {
	evenRes := 0
	for _, num := range arr {
		if num%2 == 0 {
			evenRes += num
		}
	}
	return evenRes
}
func OddSum(arr []int) int {
	oddRes := 0
	for _, num := range arr {
		if num%2 != 0 {
			oddRes += num
		}
	}
	return oddRes
}
