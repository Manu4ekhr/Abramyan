package case1_20

import "fmt"

func Case3(month int) {
	switch month {
	case 1, 2, 12:
		fmt.Println( /*Декабрь, Январь, Фераль*/ "Зима")
	case 3, 4, 5:
		fmt.Println( /*Март, Апрель, Май*/ "Весна")
	case 6, 7, 8:
		fmt.Println( /*Июнь, Июль, Август*/ "Лето")
	case 9, 10, 11:
		fmt.Println( /*Сентябрь, Октябрь, Ноябрь*/ "Осень")
	default:
		fmt.Println("Ошибка")
	}
}
