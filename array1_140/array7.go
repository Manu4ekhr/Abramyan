package array1_140

import "fmt"

func Array7() {
	n := 10
	var numbers []int
	for i := n; i > 0; i-- {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)
}
