package main

import (
	"fmt"
)

func main() {
	personSalary := map[string]int{
		"theanh": 18,
		"ductri": 20,
	}

	// key, key_exixt := personSalary["thanh"]
	// fmt.Println(key)
	// fmt.Println(key_exixt)
	for key, value := range personSalary {
		fmt.Println("key la: %s, value: %d", key, value)
	}
}
