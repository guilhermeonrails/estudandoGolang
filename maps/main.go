package main

import "fmt"

func main() {
	meuMap := make(map[int]string)
	meuMap[0] = "Guilherme"
	meuMap[1] = "Lima"
	fmt.Println(meuMap)

	outroMap := make(map[int]int)
	outroMap[0] = 42
	fmt.Println(outroMap[0])

	ultimoMap := make(map[string]string)
	ultimoMap["0"] = "Guilherme"
	fmt.Println(ultimoMap)
}
