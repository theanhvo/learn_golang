package rectangle

import "math"
import "fmt"

/*
 * Hàm init được thêm vào
 */
func init() {
	fmt.Println("rectangle package initialized")
}
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
