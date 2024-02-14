package for1_40

import (
	"fmt"
	"math"
)

func For11(n int) {
	sumSqr := 0
	for i := n; i <= n; i++ {
		sumSqr += int(math.Pow(float64(n+i), 2))
	}
	fmt.Println(sumSqr)
}
