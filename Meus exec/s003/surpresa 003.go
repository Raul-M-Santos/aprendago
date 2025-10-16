package main

import "fmt"

func main() {
	for a := 33; a <= 122; a++ {
		fmt.Printf("%#U \n", a)
	}
}

