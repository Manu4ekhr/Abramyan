package if1_30

import "fmt"

func If5(a, b, c int) {
	x := 0
	y := 0
	if a > 0 {
		x++
	} else if a < 0 {
		y++
	}
	if b > 0 {
		x++
	} else if b < 0 {
		y++
	}
	if c > 0 {
		x++
	} else if c < 0 {
		y++
	}
	fmt.Println("Количество положительных чисел:", x)
	fmt.Println("Количество отрицательных чисел:", y)
}
