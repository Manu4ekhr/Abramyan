package case1_20

import "fmt"

func Case5(n int, a, b float64) {
	switch n {
	case 1:
		fmt.Printf("%.2f", a+b)
	case 2:
		fmt.Printf("%.2f", a-b)
	case 3:
		fmt.Printf("%.2f", a*b)
	case 4:
		fmt.Printf("%.2f", a/b)
	default:
		fmt.Printf("Ошибка")
	}
}
