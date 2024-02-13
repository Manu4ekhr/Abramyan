package for1_40

import "fmt"

func For8(a, b int) {
	multiplier := 1
	for i := a; i <= b; i++ {
		multiplier = multiplier * i
	}
	fmt.Println(multiplier)
}
