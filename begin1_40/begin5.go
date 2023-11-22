package begin1_40

import "fmt"

func Begin5() {
	var a int
	fmt.Println("Вводите длину ребра куба: ")
	fmt.Scan(&a)
	v := a * a * a
	s := 6 * a * a
	fmt.Printf("Объем куба: %v \nПлощадь его поверхности: %v", v, s)
}
