package begin1_40

import (
	"fmt"
	"math"
)

func Begin20(x1, y1, x2, y2 float64) {
	distance := math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))
	fmt.Println(distance)
}
