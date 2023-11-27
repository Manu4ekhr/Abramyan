package begin1_40

import "fmt"

func Summ(a, b int) {
	result := a + b
	fmt.Println(result)
}
func Diff(a, b int) {
	result := a - b
	fmt.Println(result)
}
func Mult(a, b int) {
	result := a * b
	fmt.Println(result)
}
func Div(a, b int) {
	result := a / b
	fmt.Println(result)
}
func Abs(a int) {
	result := a
	if result < 0 {
		result *= 1
	}
	fmt.Println(result)
}

func Begin11(a, b int) {
	Summ(a, b)
	Diff(a, b)
	Mult(a, b)
	Div(a, b)
	Abs(a)
}
