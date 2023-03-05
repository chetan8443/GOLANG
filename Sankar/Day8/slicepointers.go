package main
import "fmt"
func main(){
	nums := []int {1,2,3,4,5,5}
	vari1 := &nums[0]
	fmt.Println(nums)
	*vari1 = nums[4]+nums[5]
	fmt.Println(nums)
