package main

func main() {
	//operador +
	resultadoDaSoma := 30 + 12
	println(resultadoDaSoma)
	nome := "Guilherme"
	sobrenome := "Lima"
	println(nome + sobrenome)

	//operador -
	resultadoDaSubtracao := 100 - 37
	println(resultadoDaSubtracao)

	//operador %
	restoDaDivisao := 5 % 2
	println(restoDaDivisao)

	//operador inc
	inc := 1
	inc++
	inc++
	println("O valor do inc é", inc)

	//operador dec
	dec := 10
	dec--
	dec--
	println("O valor do dec é", dec)

	//operador de atribuição
	saldo := 10
	saldo += 500
	debito := 200
	saldo -= debito
	println(saldo)
	saldo *= 2
	println(saldo)
}
