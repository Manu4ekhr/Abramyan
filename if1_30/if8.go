package if1_30

import "fmt"

func If8(a, b float64) {
	if a > b {
		fmt.Println("Большее число:", a)
		fmt.Println("Меньшее число:", b)
	}
}
