package boolean1_40

import "fmt"

func Boolean13(a, b, c int) {
	x := a < 0 || b > 0 || c > 0
	fmt.Println(x)
}
