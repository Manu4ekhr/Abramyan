package if1_30

import "fmt"

func If3(a int) {
	if a > 0 {
		fmt.Println(a + 1)
	} else if a < 0 {
		fmt.Println(a + 2)
	} else if a == 0 {
		fmt.Println(10)
	}
}
