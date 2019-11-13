package main

import (
	"fmt"
	"reflect"
)

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// for i, v := range pow {
	// 	fmt.Printf("i = %d, v = %d \n", i, v)
	// }
	for _, v := range pow {
		fmt.Printf("v = %d \n", v)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// dict := map[string]string{"name": "tri", "old": "26"}

	type_map := map[string]interface{}{"name": "theanh", "age": 2}
	fmt.Println(reflect.TypeOf(type_map))

}
