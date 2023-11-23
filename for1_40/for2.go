package for1_40

import "fmt"

func For2(a, b int) {
	//k := 0
	for i := a; i <= b; i++ {
		fmt.Println(i)
		//k++
	}
	fmt.Println(b - a + 1)
}
