package main

import (
	"fmt"
)

func tinhTong(a int, b int) int {
	return a + b
}
func main() {
	a, b := 1, 2
	total := tinhTong(a, b)
	fmt.Println(total)
}
