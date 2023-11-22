package array1_140

import "fmt"

func Array2() {
	n := 10
	var numbers []int
	numbers = append(numbers, 2)
	for i := 1; i < n; i++ {
		numbers = append(numbers, 2*numbers[i-1])
	}
	fmt.Println(numbers)
}
