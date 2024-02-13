package for1_40

import "fmt"

func For9(a, b int) {
	sumSquare := 0
	for i := a; i <= b; i++ {
		sumSquare += i * i
	}
	fmt.Println(sumSquare)
}
