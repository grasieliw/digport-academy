package main

import "fmt"

func main() {
	//variaveis
	var num1 int
	var num2 int
	var total int

	num1 = 2

	//perguntar ao usuario um numero
	fmt.Println("Olá, qual numero voce gostaria somar?")

	//ler o valor do usuario
	fmt.Scan(&num2)

	//somar os dois numeros
	total = num1 + num2

	// printar no terminal o valor total
	fmt.Printf("o valor total eh %d", total)

}
