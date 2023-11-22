package begin1_40

import "fmt"

const pi = 3.14

func Begin14(l float64, r float64, s float64) {
	l = 2 * pi * r
	s = pi * r * r
	fmt.Println("Введите длину окружности: ")
	fmt.Println("Введите значение для радиуса: ")
	fmt.Println("Введите значение для площади: ")
	fmt.Scan(&l)
	fmt.Scan(&r)
	fmt.Scan(&s)
}
