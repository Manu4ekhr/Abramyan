package for1_40

import (
	"fmt"
	"math"
)

func For13(n int) {
	result := 0.0
	for i := 1; i <= n; i++ {
		result += math.Pow(-1, float64(i+1)) * (1.0 + float64(i)/10)
	}
	fmt.Println(result)
}
