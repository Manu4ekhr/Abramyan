package if1_30

import "fmt"

func If11(a, b int) {
	if a > b {
		b = a
	} else if a < b {
		a = b
	} else {
		a = 0
		b = 0
	}
	fmt.Println(a, b)
}
