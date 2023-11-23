package begin1_40

import "fmt"

const Pi = 3.14

func Begin14(r float64) {
	l := 2 * Pi * r
	s := Pi * r * r
	fmt.Printf("Длина окружности: %v \nПлощадь окружности: %v", l, s)
}
