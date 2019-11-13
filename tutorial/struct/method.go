package main

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
 phương thức displaySalary() có vật nhận là Employee
*/

func (e Employee) displaySalary() {
	fmt.Println("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //gọi phương thức displaySalary() của kiểu Employee
}
