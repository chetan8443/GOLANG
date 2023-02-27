package main
import (
	"fmt"
	"sort"
)
func main() {
	areas := []string{"Hyderabad", "Coimbatore","Vijayawada","Bangalore"}
	sort.Strings(areas)

	fmt.Println(areas)
}
