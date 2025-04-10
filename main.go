package main

import (
	"fmt"
	
)

func main() {
	var numero [5]int
	var soma int
	for i := 0; i <5;i++ {
	fmt.Println("digite cinco numeros:" )
	fmt.Scan(&numero[i])
	soma += numero[i]
}


fmt.Println("A soma dos números é:", soma)
}
	
	
