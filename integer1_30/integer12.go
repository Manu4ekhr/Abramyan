package integer1_30

import "fmt"

func Integer12(n int) {
	a := n / 100
	b := (n / 10) % 10
	c := n % 10
	m := 100*c + 10*b + a
	fmt.Println(m)
}
