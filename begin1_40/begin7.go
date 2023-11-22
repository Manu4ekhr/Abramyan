package begin1_40

import "fmt"

func Begin7() {
	var r float64
	fmt.Println("Введите радиус круга: ")
	fmt.Scan(&r)
	pi := 3.14
	l := 2 * pi * r
	s := pi * r * r
	fmt.Printf("Длина окружности: %v \nПлощадь круга: %v", l, s)
}
