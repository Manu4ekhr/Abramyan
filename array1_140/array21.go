package array1_140

import "fmt"

func Array21() {
	var n, k, l int
	fmt.Scan(&n, &k, &l)
	var numbers []int
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}
	summ := 0.0
	count := 0
	for i := k - 1; i <= l-1; i++ {
		summ += float64(numbers[i])
		count++
	}
	middle := (summ / float64(count))
	fmt.Println(middle)
}
