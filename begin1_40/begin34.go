package begin1_40

import "fmt"

func Begin34(x, ar, y, br float64) {
	chococandies := ar / x
	fmt.Printf("Стоимость кг шоколадной конфеты в рублях: %v \n", chococandies)
	toffee := br / y
	fmt.Printf("Стоимость кг ирисок в рублях: %v \n", toffee)
	moreexp := chococandies / toffee
	fmt.Printf("Шоколадные конфеты дороже ирисок в: %v \n", moreexp)
}
