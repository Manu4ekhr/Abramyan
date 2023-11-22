package begin1_40

import (
	"fmt"
	"math"
)

func Hypotenuse(a, b float64) float64 {
	c := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
	fmt.Println(c)
	return c
}
func Perimeter(a, b, c float64) {
	p := a + b + c
	fmt.Println(p)
}
func Task12(a, b float64) {
	c := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
	fmt.Println(c)

	p := a + b + c
	fmt.Println(p)
}
