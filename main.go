package main

import "fmt"

func main() {
	var usuario string
	var senha string
	fmt.Println("digite um nome de usuario e uma senha")
	fmt.Scan(&usuario)
	fmt.Scan(&senha)
	if usuario == "admin" && senha == "1234"{
		fmt.Println("usuario correto")
	}else {
		fmt.Println("usuario incorreto")
	}
}
