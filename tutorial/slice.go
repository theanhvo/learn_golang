package main

import (
	"fmt"
	// "reflect"
)

func main() {
	// a := [5]int{76, 77, 78, 79, 80}
	// var b []int = a[1:4] //creates a slice from a[1] to a[3]
	// fmt.Println(b)
	// fmt.Println(reflect.TypeOf(b))
	// c := []int{6, 7, 8} //creates and array and returns a slice reference
	// fmt.Println(c)

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		fmt.Println(i)
		dslice[i]++
	}
	fmt.Println("array after", darr)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6

}
