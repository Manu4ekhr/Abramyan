package case1_20

import "fmt"

func Case2(K int) {
	switch K {
	case 1:
		fmt.Println("Плохо")
	case 2:
		fmt.Println("Неудовлетворительно")
	case 3:
		fmt.Println("Удовлетворительно")
	case 4:
		fmt.Println("Хорошо")
	case 5:
		fmt.Println("Отлично")
	default:
		fmt.Println("Ошибка")
	}
}
