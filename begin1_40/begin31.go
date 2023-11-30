package begin1_40

import "fmt"

func Begin31(Tf float64) {
	var Tc float64
	Tc = (Tf - 32) * 5 / 9
	fmt.Printf("Значение температуры в градусах Цельсия: %v \n", Tc)
}
