package case1_20

import "fmt"

func Case4(month int) {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(31)
	case 4, 6, 9, 11:
		fmt.Println(30)
	case 2:
		fmt.Println(28)
	default:
		fmt.Println("Ошибка")
	}
}
