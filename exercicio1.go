package main

import "fmt"

func main() {
	//variaveis
	var name string

	//perguntar ao usuario qual o seu nome
	fmt.Println("Olá, qual o seu nome?")

	//escrever no terminal
	fmt.Scan(&name)

	// printar no terminal o nome
	fmt.Printf("Bem vindo %s", name)

}
