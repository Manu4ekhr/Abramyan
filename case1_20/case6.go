package case1_20

import "fmt"

func Case6(a int, b float64) {
	switch a {
	case 1:
		fmt.Println(b * 0.1)
	case 2:
		fmt.Println(b * 1000)
	case 3:
		fmt.Println(b * 1)
	case 4:
		fmt.Println(b * 0.001)
	case 5:
		fmt.Println(b * 0.01)
	default:
		fmt.Println("Ошибка")
	}
}
