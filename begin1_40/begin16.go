package begin1_40

import (
	"fmt"
	"math"
)

func Begin16(x1, x2 float64) {
	var d float64
	d = math.Abs(x1) - math.Abs(x2)
	fmt.Println(d)
}
