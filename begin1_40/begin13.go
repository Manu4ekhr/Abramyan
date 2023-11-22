package begin1_40

import (
	"fmt"
	"math"
)

func Begin13(r1, r2 float64) {
	s1 := math.Pi * r1 * r1
	s2 := math.Pi * r2 * r2

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 - s2)
}
