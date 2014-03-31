package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Print("Escreva", i, "como ")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ëh final de semana")
	default:
		fmt.Println("Ëh dia de semana")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("É antes de meio dia")
	default:
		fmt.Println("E depois do meio dia")
	}
}
