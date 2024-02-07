package if1_30

import "fmt"

func If10(a, b int) {
	if a != b {
		a = a + b
		b = a
	} else {
		a, b = 0, 0
	}
	fmt.Println(a, b)
}
