package main

import (
	"fmt"
)

func main() {
	// var luong map[string]string
	luong := make(map[string]string)
	luong["thang1"] = "1000"
	fmt.Println(luong)
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)

}
