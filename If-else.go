
package main

import "fmt"

func main()  {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	}else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 é divisivel por 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "è negativo")
	}else if num < 10  {
		fmt.Println(num, "tem 1 digito")
	}else {
		fmt.Println(num, "Tem multiplos digitos")
	}
}