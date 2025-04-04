package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	greeting := "ola mundo"
	fmt.Println(strings.Contains(greeting, "mundo"))
	fmt.Println(strings.ReplaceAll(greeting, "ola", "oi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "mundo"))
	fmt.Println(strings.Split(greeting, "mundo"))
	ages := []int {50, 80, 10}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 50)
	fmt.Println(index)
	names := []string{"murilo","maicon","fabiano"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "murilo"))
}
