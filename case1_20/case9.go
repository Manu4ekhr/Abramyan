package case1_20

import (
	"fmt"
)

func Case9(d, m int) {

	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		if d == 31 {
			d = 0
		}
	case 4, 6, 9, 11:
		if d == 30 {
			d = 0
		}
	case 2:
		if d == 28 {
			d = 0
		}
	}
	if d == 0 {
		if m == 12 {
			m = 1
		} else {
			m++
		}
		d++
	} else {
		d++
	}
	fmt.Printf("%d.%d\n", d, m)
}
