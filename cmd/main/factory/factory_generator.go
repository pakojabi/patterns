package factory

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome int
}

// functional approach

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee{
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// stuctural approach

type EmployeeFactory2 struct {
	Position string
	AnnualIncome int
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory2 {
	return &EmployeeFactory2{position, annualIncome}
}

func (f *EmployeeFactory2) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

// use

func RunFactoryGenerator() {
	fmt.Println("Functional approach")
	empGenerator := NewEmployeeFactory("Pringui", 14000)
	emp1 := empGenerator("fulano")
	fmt.Println("generated employee: ", emp1)

	fmt.Println("Structural approach")
	empFactory := NewEmployeeFactory2("Pringui", 14000)
	emp2 := empFactory.Create("fulano")
	fmt.Println("generated employee: ", emp2)
}