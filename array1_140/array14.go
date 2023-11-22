package array1_140

import "fmt"

func Array14() {
	var b, n int
	fmt.Scan(&n)
	var a []int
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		a = append(a, b)
	}
	var numbers []int
	for i := 1; i < n; i += 2 {
		numbers = append(numbers, a[i])
	}
	for i := 0; i < n; i += 2 {
		numbers = append(numbers, a[i])
	}
	fmt.Println(numbers)
}
