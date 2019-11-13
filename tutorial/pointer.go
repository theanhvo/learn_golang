package main

import (
	"fmt"
)

func modify(arr *[3]int) {
	(*arr)[0] = 90
}

func main() {
	a := [3]int{89, 90, 91}
	fmt.Println(&a)
	modify(&a)
	fmt.Println(a)
}
