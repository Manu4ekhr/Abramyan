package begin1_40

import "fmt"

func Begin23(a, b, c int) {
	var t, p int
	t = a
	p = b
	a = c
	b = t
	c = p
	fmt.Printf("Значение для А: %v \nЗначение для B: %v \nЗначение для C: %v", a, b, c)
}
