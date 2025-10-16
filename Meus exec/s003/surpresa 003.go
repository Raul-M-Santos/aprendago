package main

import "fmt"

func main() {
	for a := 33; a <= 122; a++ {
		fmt.Printf("%#U \n", a)
	}
	resultado:= calc (10, 20)
	fmt.Println(resultado)
}


func calc(a int, b int) int {
	return a + b
}