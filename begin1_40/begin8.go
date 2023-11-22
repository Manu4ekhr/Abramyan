package begin1_40

import "fmt"

func Begin8() {
	var a, b int
	fmt.Scan(&a, &b)
	evarage := float64(a+b) / 2
	fmt.Printf("Среднее арифметическое: %v", evarage)
}
