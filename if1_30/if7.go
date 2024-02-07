package if1_30

import "fmt"

func If7(a, b float32) {
	if a < b {
		fmt.Println(1)
	} else if b < a {
		fmt.Println(2)
	}
}
