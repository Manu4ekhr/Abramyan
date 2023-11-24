package begin1_40

import (
	"fmt"
	"math"
)

func Begin18(a, b, c float64) {
	ac := math.Abs(a) * math.Abs(c)
	bc := math.Abs(b) * math.Abs(c)
	fmt.Println(ac)
	fmt.Println(bc)
}
