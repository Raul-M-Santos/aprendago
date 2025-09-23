package main

import "fmt"

func main() {
	var esportefavorito string = "Badminton"

	switch esportefavorito {
	case "Badminton":
		fmt.Println("Meu esporte favorito é badminton")
    default:
		fmt.Println("Não é badminton")
	}
}
