package for1_40

import "fmt"

func For6(p float64) {
	for m := 1.2; m <= 2.0; m += 0.2 {
		fmt.Printf("%f кг = %f \n", m, p*m)
	}
}
