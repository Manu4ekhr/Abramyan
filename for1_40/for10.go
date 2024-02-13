package for1_40

import "fmt"

func For10(n int) {
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}
	fmt.Println(sum)
}
