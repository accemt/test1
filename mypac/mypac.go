package mypac

import (
	"fmt"
)

func PrintMasT(mas [10]int) {
	var i = 0
	fmt.Print("MAS: ")
	for i < 10 {
		fmt.Print(mas[i])
		i += 1
	}
	fmt.Println()
}
