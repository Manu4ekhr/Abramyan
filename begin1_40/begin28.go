package begin1_40

import "fmt"

func Begin28(a int) {
	var x int
	x = a * a
	fmt.Printf("А в степени 2: %v \n", x)
	a = x * a
	fmt.Printf("А в степени 3: %v \n", a)
	x = x * a
	fmt.Printf("А в степени 5: %v \n", x)
	a = x * x
	fmt.Printf("А в степени 10: %v \n", a)
	x = a * x
	fmt.Printf("А в степени 15: %v \n", x)
}
