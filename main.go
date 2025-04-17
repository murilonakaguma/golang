package main

import "fmt"


func dadosPessoa(nome string, idade int) (int, string) {
	var status string
	if idade >= 18 {
		status = "maior de idade"
	} else {
		status = "menor de idade"
	}
	return idade, status
}

func main() {
	idade, status := dadosPessoa("João", 20)
	fmt.Printf(" Idade: %d\nStatus: %s\n", idade, status)
}
