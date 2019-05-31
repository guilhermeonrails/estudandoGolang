package main

import "fmt"

func main() {
	//1.0
	// minhaArray := [3]int{}
	// minhaArray[0] = 42
	// minhaArray[1] = 10
	// minhaArray[2] = 1000
	// fmt.Println(minhaArray)

	//podemos reduzir este código acima, adicionando os
	//valores dentro das chaves separados por vírgula

	//2.0
	// minhaArray := [3]int{42, 10, 1000}
	// fmt.Println(minhaArray)

	//2.5
	// minhaArray := []int{10, 20, 30, 40}
	// meuSlice := minhaArray[2:3]
	// fmt.Println(meuSlice)

	// não precisamos especificar o tamanho do array
	//3.0
	// minhaArray := [...]int{42, 10, 1000}
	// fmt.Println(minhaArray)

	// //tamanho do array
	// fmt.Println("tamanho", len(minhaArray))

	//4.0 - cap e len
	//cap informa a capacidade do array
	//len quantos itens estão presentes no array

	//ass make: []TIPO, LEN, CAP
	s := make([]int, 1, 1)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Println("capacidade", cap(s), "tamanho", i)
	}
}
