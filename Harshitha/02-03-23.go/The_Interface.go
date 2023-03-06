
package main
 
import "fmt"
 
type salaryCalculator interface {
   calculateSalary() float64
   report()
}
 
type PermanentEmployee struct {
   id          int
   basicSalary float64
   commission  float64
}
 
type ContractEmployee struct {
   id          int
   basicSalary float64
}
 
func (p PermanentEmployee) calculateSalary() float64 {
   return p.basicSalary + (p.commission/100)*p.basicSalary
}
 
func (c ContractEmployee) calculateSalary() float64 {
   return c.basicSalary
}
 
func (p PermanentEmployee) report() {
   fmt.Printf("Employee ID %d earns USD %f per month \n", p.id, p.calculateSalary())
}
 
func (c ContractEmployee) report() {
   fmt.Printf("Employee ID %d earns USD %f per month \n", c.id, c.calculateSalary())
}
 
func main() {
   var calculator salaryCalculator
   calculator = PermanentEmployee{id: 1, basicSalary: 10000, commission: 20}
   calculator.report()
   calculator = ContractEmployee{id: 2, basicSalary: 5000}
   calculator.report()
}