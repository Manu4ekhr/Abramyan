package array1_140

import "fmt"

func Array3() {
	var n, a, d int
	fmt.Println("Вводите число N: ")
	fmt.Scan(&n)
	fmt.Println("Вводите первый член арифметической прогресси А: ")
	fmt.Scan(&a)
	fmt.Println("Вводите разницу арифметической прогресси D: ")
	fmt.Scan(&d)
	var members []int
	for i := 0; i < n; i++ {
		members = append(members, a+i*d)
	}
	fmt.Println(members)
}
