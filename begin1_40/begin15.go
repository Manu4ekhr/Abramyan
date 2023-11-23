package begin1_40

import (
	"fmt"
	"math"
)

func Begin15(s float64) {
	r := math.Sqrt(s / Pi)
	d := 2 * r
	l := 2 * Pi * r
	fmt.Printf("Диаметр окружности: %v \nДлина окружности: %v", d, l)
}
