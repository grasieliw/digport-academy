package main

import "fmt"

func main() {
	var contador int

	for contador = 10; contador != 0; contador-- {
		if contador > 0 {
			fmt.Printf("Valor de contador eh %d \n", contador)
		}
	}
	fmt.Printf("Arrasou!")
}
