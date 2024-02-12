package case1_20

import (
	"fmt"
)

func Case10(n int, c rune) {

	fmt.Print("Направление:")
	_, err := fmt.Scanf("%c", &c)
	if err != nil {
		fmt.Println("Error reading character:", err)
		return
	}
	fmt.Print("Команда:")
	_, err = fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error reading number:", err)
		return
	}
	switch c {
	case 'С':
		switch n {
		case 1:
			c = 'З'
		case 0:
			c = 'С'
		case -1:
			c = 'В'
		}
	case 'В':
		switch n {
		case 1:
			c = 'С'
		case 0:
			c = 'В'
		case -1:
			c = 'Ю'
		}
	case 'Ю':
		switch n {
		case 1:
			c = 'В'
		case 0:
			c = 'Ю'
		case -1:
			c = 'З'
		}
	case 'З':
		switch n {
		case 1:
			c = 'Ю'
		case 0:
			c = 'З'
		case -1:
			c = 'С'
		}
	}
	fmt.Printf("%c\n", c)
}
