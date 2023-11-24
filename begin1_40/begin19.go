package begin1_40

import (
	"fmt"
	"math"
)

func Begin19(x1, x2, y1, y2 float64) {
	p := math.Abs(float64(x1-x2)) * math.Abs(float64(y1-y2))
	s := math.Abs(float64((x1 - x2)) * math.Abs(y1-y2))
	fmt.Println(p, s)
}
