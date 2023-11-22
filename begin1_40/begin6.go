package begin1_40

import "fmt"

func Begin6() {
	var a, b, c int
	fmt.Println("Введите длины ребра: ")
	fmt.Scan(&a, &b, &c)
	v := a * b * c
	s := 2 * (a*b + b*c + a*c)
	fmt.Printf("Объем прямоугольного параллелепипеда: %v \nПлощадь поверхности прямоугольного параллелепипеда: %v", v, s)
}
