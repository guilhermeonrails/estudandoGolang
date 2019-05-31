package main

const (
	nome      = "Guilherme"
	sobrenome = "Lima"
)

const (
	primeira = iota
	segunda
	terceira
)

const (
	quarta = iota
)

const (
	potenciaUm = 1 << (1 * iota)
	potenciaDois
	potenciaTres
)

func main() {
	println(nome, sobrenome)
	println(primeira)
	println(segunda)
	println(terceira)
	println(quarta)
	println(potenciaUm)
	println(potenciaDois)
	println(potenciaTres)
}
