package main

import "fmt"

func main() {
	var meuNumeroInteiro = 42
	fmt.Println("inteiro:", meuNumeroInteiro)

	var meuFloat float32 = 42.398392
	fmt.Println("float:", meuFloat)

	minhaString := "Esta é minha string"
	fmt.Println("String:", minhaString)

	meuNumeroComplex := complex(2, 3)
	fmt.Println("complex:", meuNumeroComplex)
	fmt.Println("complex real:", real(meuNumeroComplex))
	fmt.Println("complex imaginário:", imag(meuNumeroComplex))
}
