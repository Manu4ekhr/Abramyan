package array1_140

import "fmt"

func Array4() {
	var n, a, d int
	fmt.Println("Вводите число N: ")
	fmt.Scan(&n)
	fmt.Println("Вводите первый член геометрической прогресси А: ")
	fmt.Scan(&a)
	fmt.Println("Вводите разницу геометрической прогресси D: ")
	fmt.Scan(&d)
	var members []int
	members = append(members, a)
	for i := 1; i < n; i++ {
		members = append(members, members[i-1]*d)
	}
	fmt.Println(members)
}
