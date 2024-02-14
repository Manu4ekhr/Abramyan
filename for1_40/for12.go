package for1_40

import "fmt"

func For12(n int) {
	result := 1.0
	for i := 0; i <= n; i++ {
		result *= 1.0 + float64(i)/10
	}
	fmt.Println(result)
}
