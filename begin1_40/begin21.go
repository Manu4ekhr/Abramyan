package begin1_40

import (
	"fmt"
	"math"
)

func Begin21(x1, y1, x2, y2, x3, y3 float64) {
	A := math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))
	B := math.Sqrt(((x3 - x2) * (x3 - x2)) + ((y3 - y2) * (y3 - y2)))
	C := math.Sqrt(((x3 - x1) * (x3 - x1)) + ((y3 - y1) * (y3 - y1)))
	p := (A + B + C) / 2
	s := math.Sqrt(p * (p - A) * (p - B) * (p - C))
	fmt.Printf("Площадь треугольника: %v \nПериметр треугольника: %v", s, p)
}

//fmt.Printf("Диаметр окружности: %v \nДлина окружности: %v", d, l)
