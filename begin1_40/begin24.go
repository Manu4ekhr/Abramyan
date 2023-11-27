package begin1_40

import "fmt"

func Begin24(a, b, c int) {
	a = b + c - 6  // (a = c-8); a = 6 + 8 - 6 = 8
	b = a + c - 12 // (b = a-4); b = 8 + 8 - 12 = 4
	c = b + a - 6  // (c = b-6); c = 4 + 8 - 6 = 4
	fmt.Printf("Значение для A: %v \nЗначение для B: %v \nЗначение для C: %v", a, b, c)
}
