package main

import (
	"fmt"
	
)

func main() {
	age := 16
	fmt.Println(age <= 20)
	fmt.Println(age >= 20)
	fmt.Println(age == 20)
	fmt.Println(age != 20)
	if age < 20 {
		fmt.Println("menor que 20 anos")
	} else if age < 10 {
		fmt.Println("menor que 10 anos")
	} else {
		fmt.Println("nao é menor que 20 anos")
	}
	names := []string{"manu", "lucas", "bruce","murilo","robin","igor"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continue apos a posiçao", index, "e valor", value)
			continue
		}
		if index > 2 {
			fmt.Println("sair apos", index)
			break
		}
		fmt.Println("valor:", index, value)
	}
}
	
	
