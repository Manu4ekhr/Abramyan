package begin1_40

import "fmt"

func Begin32(Tc float64) {
	var Tf float64
	Tf = (Tc * 9 / 5) + 32
	fmt.Printf("Значение температуры в градусах Фаренгейта: %v \n", Tf)
}
