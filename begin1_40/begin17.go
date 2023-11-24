package begin1_40

import (
	"fmt"
	"math"
)

func Begin17(a, b, c float64) {
	ac := math.Abs(c) - math.Abs(a)
	bc := math.Abs(c) - math.Abs(b)
	k := ac + bc
	fmt.Println(ac)
	fmt.Println(bc)
	fmt.Println(k)
}
