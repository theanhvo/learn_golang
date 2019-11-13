package main

import "fmt"

func main() {
	// make(map[type of key]type of value
	languages := make(map[string]float32)

	// add item enter map
	languages["python"] = 0.2
	languages["java"] = 0.3

	fmt.Println(languages)
}
