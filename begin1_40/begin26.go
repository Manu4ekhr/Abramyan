package begin1_40

import (
	"fmt"
	"math"
)

func Begin26(x float64) {
	y := 4*math.Pow((x-3), 6) - 7*math.Pow((x-3), 3) + 2
	fmt.Printf("Значение функции: %v \n", y)
}
