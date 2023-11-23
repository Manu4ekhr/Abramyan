package for1_40

import "fmt"

func For1(k, n int) {
	fmt.Println("Введите значение для K:")
	fmt.Println("Введите значение для N:")
	fmt.Scan(&k)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Println(k)
	}
}
