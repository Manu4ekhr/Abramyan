package integer1_30

import "fmt"

func Integer11(n int) {
	a := n / 100
	b := (n / 10) % 10
	c := n % 10
	fmt.Println(a, b, c)
}
