package begin1_40

import "fmt"

func Begin22(a, b int) {
	a = a + b
	b = a - b
	a = a - b
	//c = a
	//a = b
	//b = c
	fmt.Printf("Значение для А: %v \nЗначение для B: %v", a, b)
}

/*
a = a + b
b = a - b
a = a - b
*/
