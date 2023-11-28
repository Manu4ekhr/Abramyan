package begin1_40

import "fmt"

func Begin28(a int) {
	var x, y int
	x = a * a
	fmt.Printf("А в степени 2: %v \n", x)
	y = a * x
	fmt.Printf("А в степени 3: %v \n", y)
	y = y * x
	fmt.Printf("А в степени 5: %v \n", y)
	y = y * y
	fmt.Printf("А в степени 10: %v \n", y)
	y = y * x * a
	fmt.Printf("А в степени 15: %v \n", y)
}
