package main

import (
	"fmt"
)

func main() {
	// a := 100
	// var pointer *int
	// pointer = &a
	// fmt.Printf("%T", pointer)

	// b := 123
	// p2 := new(int)
	// p2 = &b
	// fmt.Println(p2)

	// zero value
	// var pointer *int
	// fmt.Println(pointer)
	// pointer2 := new(int)
	// fmt.Println(pointer2)

	// var pointer *int
	// a := 100
	// pointer = &a
	// fmt.Println(a)

	// *pointer = 999 // a := 999
	// fmt.Println(pointer)
	// fmt.Println(a)

	// pointer arr
	array := [3]int{1, 2, 3}
	var pointer *[3]int
	pointer = &array
	fmt.Println(pointer)

	c := 30
	var pointer2 *int = &c
	ApplyPointer(pointer2)
	fmt.Println(c)

}

func ApplyPointer(pointer *int) {
	*pointer = 777
	fmt.Println(pointer)
}
