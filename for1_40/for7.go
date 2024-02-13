package for1_40

import "fmt"

func For7(a, b int) {
	summ := 0
	for i := a; i <= b; i++ {
		summ += i
	}
	fmt.Println(summ)
}
