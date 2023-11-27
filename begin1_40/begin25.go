package begin1_40

import (
	"fmt"
	"math"
)

func Begin25(x float64) {
	y := 3*math.Pow(x, 6) - 6*math.Pow(x, 2) - 7
	fmt.Printf("Значение функции: %v \n", y)
}
