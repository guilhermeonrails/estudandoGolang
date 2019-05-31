package main

import "fmt"

func main() {
	saudacoes("Olá", "Saudações", "Oi")
}

func saudacoes(comprimentos ...string) {
	for _, comprimento := range comprimentos {
		fmt.Println(comprimento)
	}
}
