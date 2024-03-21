package main

import "fmt"

func main() {
	//variaveis
	var numero int

	//perguntar ao usuario um numero inteiro
	fmt.Println("Digite um numero inteiro")

	//ler o valor do numero inteiro
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Printf("Positivo!")
	} else if numero < 0 {
		fmt.Printf("Negativo!")
	} else {
		fmt.Printf("Zero!")
	}
}
