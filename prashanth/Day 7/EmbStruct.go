import "fmt"


type Department struct {
	name string
	age  int
}

type Employee struct {
	  Person
	id    int
	title string
}

func main() {

	emp := Employee{
		Person: Person{
			name: "Ravikumar",
			age:  35,
		},
		id:    110,
		title: "Orthopedician",
	}

	
	fmt.Println(emp.name) 
	fmt.Println(emp.age)  

	fmt.Println(emp.id)    
	fmt.Println(emp.title)
}
