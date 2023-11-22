package begin1_40

import (
	"fmt"
)

func Begin10(a, b int) {
	summ := a*a + b*b
	diff := a*a - b*b
	comp := a * a * b * b
	quot := a * a / b * b
	fmt.Println("Сумма:", summ, "Разность:", diff, "Произведение:", comp, "Частное:", quot)
}
