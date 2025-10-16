package main 

import (
	"fmt"
)

func main () {
	resul := calc (10, 1, "-")
	fmt.Println(resul)
}

func calc (a int, b int, c string) int {
	switch {
	case c == "+":
		return a + b
	case c == "-":
		return a - b
	case c == "*":
		return a * b
	case c == "/":
		if b == 0 {
			fmt.Println("Não pode divisão por 0")
		}else {
			return a / b
		}
		default :
		fmt.Println("digite um operador valido")
	}
	return 0
}