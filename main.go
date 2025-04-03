package main

import "fmt"

func main() {
	var ages = [4]int{17, 16, 20 ,21}
	nomes := [4]string{"murilo","jaspion","lucas","bugio"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[2] = "gobetti"
	fmt.Println(nomes)
}
