package array1_140

import "fmt"

func Array1() {
	N := 10
	var numbers []int
	for i := 1; i <= N; i++ {
		numbers = append(numbers, 2*i-1)
	}
	fmt.Println(numbers)
}

/* func Array1_a() {
	N := 10
	var numbers []int
	numbers = append(numbers, 1)
	for i := 1; i < N; i++ {
		numbers = append(numbers, numbers[i-1]+2)
	}
	fmt.Println(numbers)
}*/
