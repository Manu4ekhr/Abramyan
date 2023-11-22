package if1_30

import "fmt"

func If4(a, b, c int) {
	n := 0
	if a > 0 {
		n++
	}
	if b > 0 {
		n++
	}
	if c > 0 {
		n++
	}
	fmt.Println(n)
}
