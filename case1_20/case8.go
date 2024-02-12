package case1_20

import "fmt"

func Case8(D, M int) {

	D = D - 1
	if D == 0 {
		M = M - 1
		if M == 0 {
			M = 12
		}
		switch M {
		case 1, 3, 5, 7, 8, 10, 12:
			D = 31
		case 4, 6, 9, 11:
			D = 30
		case 2:
			D = 28
		}
	}
	fmt.Printf("Предыдущая дата:%d.%d\n", D, M)
}
