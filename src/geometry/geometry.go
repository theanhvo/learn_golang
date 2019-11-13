package main

import (
	"fmt"
	"geometry/rectangle" //importing package tùy chỉnh
	"log"
)

/*
 * 1. Biến toàn cục (package variables)
 */
var rectLen, rectWidth float64 = 6, 7

/*
*2. Hàm init để kiểm tra nếu chiều rộng và chiều dài nhỏ hơn 0
 */
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
