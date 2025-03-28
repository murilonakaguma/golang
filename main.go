package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int
	fmt.Println("digite um numero:")
	fmt.Scan(&numero1)
	fmt.Println("digite um numero:")
	fmt.Scan(&numero2)
	fmt.Println("a soma dos numeros é:", numero1 + numero2)
	fmt.Println("a subtraçao do numeros é:", numero1 - numero2)
	fmt.Println("a multiplicaçao do numero é:", numero1 * numero2)
	fmt.Println("a divisao do numero é:", numero1 / numero2)
	fmt.Println("o resto da divisao é:", numero1 % numero2)
}
