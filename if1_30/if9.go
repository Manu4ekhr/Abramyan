package if1_30

import "fmt"

func If9(a, b float32) {
	if a > b {
		a = a + b
		b = a - b
		a = a - b
	}
	fmt.Println("Новое значение A:", a)
	fmt.Println("Новое значение В:", b)
}
