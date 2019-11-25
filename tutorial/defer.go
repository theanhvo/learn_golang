package main

import (
	"fmt"
)

func Hello() {
	fmt.Println("Test")
	defer fmt.Println("World")
	fmt.Println("Hello")
	fmt.Println("The")
}
func main() {
	fmt.Println("say")
	Hello()
}
