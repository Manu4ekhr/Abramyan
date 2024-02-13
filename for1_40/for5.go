package for1_40

import "fmt"

func For5(p float64) {
	//cost = price * mass(or weight)

	for m := 0.1; m < 1.0; m += 0.1 {
		fmt.Printf("%f kg = %f \n", m, p*m)
	}
}
