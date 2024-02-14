package for1_40

import "fmt"

func For14(n int) {
	result := 0
	for i := 1; i <= (2*n - 1); i += 2 {
		result += i
	}
	fmt.Println(result)
}
