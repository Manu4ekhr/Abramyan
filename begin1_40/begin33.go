package begin1_40

import "fmt"

func Begin33(k, p, y float64) {
	cost1 := p / k
	costY := p / k * y
	fmt.Printf("Стоимость 1кг вкусной конфеты в рублях: %v \nСтоимость Yкг вкусной конфеты в рублях: %v \n", cost1, costY)
}

/*package main

import "fmt"

func main() {
	var X, A, Y, B float32
	fmt.Print("X:")
	fmt.Scan(&X)
	fmt.Print("A:")
	fmt.Scan(&A)
	fmt.Print("Y:")
	fmt.Scan(&Y)
	fmt.Print("B:")
	fmt.Scan(&B)
	var ChocolatePrice float32
	ChocolatePrice = A / X
	fmt.Printf("Цена шоколадных конфет:%f\n", ChocolatePrice)
	var ToffeesPrice float32
	ToffeesPrice = B / Y
	fmt.Printf("Цена ирисок:%f\n", ToffeesPrice)
	fmt.Printf("Шоколад дороже ирисок в %f раза\n", ChocolatePrice/ToffeesPrice)
}*/
