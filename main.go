package main

import "fmt"

func main() {
	var ages = [4]int{17, 16, 20 ,21}
	nomes := [4]string{"murilo","jaspion","lucas","bugio"}
	fmt.Println(ages)
	fmt.Println(nomes)
	nomes[2] = "gobetti"
	fmt.Println(nomes)
	// Slice 
	var Score = []int{100, 200, 300,400}
	fmt.Println(Score)
	Score[1] = 2
	fmt.Println(Score, len(Score), cap(Score))
	rangeOne := Score[1:3]
	fmt.Println(rangeOne)
	rangeTwo := Score[2:]
	fmt.Println(rangeTwo)
	rangeThree := Score[:3]
	fmt.Println(rangeThree)
	var superherois = []string{"robin","flash","iron man"}
	fmt.Println(superherois)
	superherois = append(superherois, "ben 10", "batman")
	fmt.Println(superherois, len(superherois), cap(superherois))
	superherois[4] = "robin"

}
