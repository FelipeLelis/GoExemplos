
package main

import (
	"fmt"
	"time"
)

func main()  {
	i := 2
	fmt.Println("Escreva ", i, " é ")
	switch i {
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
	case 3:
		fmt.Println("Tres")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("É final de semana")
	default:
		fmt.Println("É dia de semana")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("È de manha")
	default:
		fmt.Println("Já é de tarde")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Eu sou um booleano")
		case int:
			fmt.Println("Eu sou um inteiro")
		default:
			fmt.Println("Eu não sei o tipo", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}