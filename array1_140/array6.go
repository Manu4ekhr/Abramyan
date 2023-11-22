package array1_140

import "fmt"

func Array6() {
	var n, a, b int
	n = 10
	a = 1
	b = 3

	sum := a + b
	var f []int

	f = append(f, a, b)

	for i := 2; i < n; i++ {
		f = append(f, sum)
		sum = sum + f[i]
	}
	fmt.Println(f)
}
