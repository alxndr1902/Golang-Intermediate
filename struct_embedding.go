package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person
	employeeId string
	salary     float64
}

func (p person) introduce() {
	fmt.Printf("Hi, my name is %s, and I'm %d years old\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, and I earn %.2f\n", e.name, e.salary)
}

func main() {
	emp := Employee{
		person: person{
			name: "John",
			age:  40,
		},
		employeeId: "123",
		salary:     1000,
	}

	fmt.Println(emp.name)
	fmt.Println(emp.age)
	fmt.Println(emp.employeeId)
	fmt.Println(emp.salary)

	emp.introduce()
}
