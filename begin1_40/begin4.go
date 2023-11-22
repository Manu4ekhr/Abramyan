package begin1_40

import (
	"fmt"
)

func Begin4() {
	var d int
	fmt.Println("Введите диаметр окружности: ")
	fmt.Scan(&d)
	pi := 3.14
	result := pi * float64(d)
	fmt.Println(result)
}
