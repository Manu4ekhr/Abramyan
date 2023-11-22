package array1_140

import "fmt"

func Array5() {
	n := 10
	var fib []int
	f0 := 1
	f1 := 1
	fib = append(fib, f0, f1)
	for i := 2; i < n; i++ {
		fib = append(fib, f1+f0)
		f0 = f1
		f1 = fib[i]
	}
	fmt.Println(fib)
}
