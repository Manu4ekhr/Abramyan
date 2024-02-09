package if1_30

import "fmt"

func If12(a, b, c float32) {
	if a < b && a < c {
		fmt.Println(a)
	} else if b < a && b < c {
		fmt.Println(b)
	} else if c < a && c < b {
		fmt.Println(c)
	}
}
