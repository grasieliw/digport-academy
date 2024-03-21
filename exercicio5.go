package main

import "fmt"

var a, b, c, total int

func Soma(a int, b int, c int) int {
	switch c {
	case 1:
		total = a + b
	case 2:
		total = a - b
	case 3:
		total = a * b
	case 4:
		total = (a / b)
	}
	return total
}

func main() {
	fmt.Println("Digite um numero inteiro positivo")
	fmt.Scan(&a)

	fmt.Println("Agora, digite outro numero inteiro positivo")
	fmt.Scan(&b)

	fmt.Println("Escolha uma operacao: 1=SOMA  2= SUBTRACAO  3=MULTIPLICACAO  4=DIVISAO")
	fmt.Scan(&c)

	ValorTotal := Soma(a, b, c)
	fmt.Printf("o valor total eh %d", ValorTotal)
}
