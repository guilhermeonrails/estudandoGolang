package main

import "fmt"

func main() {
	// meuArray := []int{42, 1000, 17, 40}
	// meuSlice := meuArray[:]
	// fmt.Println(meuSlice)
	// fmt.Println(len(meuSlice))

	// meuSlice = append(meuSlice, 12345)
	// fmt.Println(meuSlice)
	// fmt.Println(len(meuSlice))

	//outra forma criar slices
	// meuSlice := []float32{42, 19}
	// fmt.Println(meuSlice)
	// fmt.Println(len(meuSlice))

	//ass make: []TIPO, LEN, CAP
	meuSlice := make([]int, 5)
	meuSlice[5] = 42
	fmt.Println(meuSlice[5])
}
