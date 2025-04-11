package main

import (
	"fmt"
	
)
func sayGreeting(nome string){
	fmt.Println("olá!", nome)
}
func addNumber(numero1 int, numero2 int) int{
	return numero1 + numero2
}
func main() {
	sayGreeting("juvelino")
	resultado := addNumber(10,20)
	fmt.Println(resultado)

	}