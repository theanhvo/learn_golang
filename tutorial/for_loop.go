package main

import "fmt"

var test = "er"

func main() {
	i := 1
	for i <= 3 {
		// fmt.Println(i)
		i = i + 1
	}

	// for range
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println(test)

	for item := 0; item <= 30; item++ {
		if item%3 == 0 && item%5 == 0 {
			fmt.Println(item)
		}
	}
}
