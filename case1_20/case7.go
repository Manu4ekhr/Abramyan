package case1_20

import "fmt"

func Case7(a int, b float32) {
	switch a {
	case 1:
		fmt.Println(b * 1)
	case 2:
		fmt.Println(b * 1e-6)
	case 3:
		fmt.Println(b * 0.001)
	case 4:
		fmt.Println(b * 1000)
	case 5:
		fmt.Printf("%.1f", b*100)
	default:
		fmt.Println("Ошибка")
	}
}
