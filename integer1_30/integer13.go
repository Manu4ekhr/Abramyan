package integer1_30

import "fmt"

func Integer13(n int) {
	a := n / 100
	b := (n / 10) % 10
	c := n % 10
	m := 100*b + 10*c + a
	fmt.Println(m)
}
