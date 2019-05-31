package main

import "fmt"

type pessoa struct {
	nome  string
	saldo float64
}

func (p pessoa) getSaldo() float64 {
	return p.saldo
}
func (p pessoa) getNome() string {
	return p.nome
}

func (p pessoa) depositar(valor float64) float64 {
	return p.saldo + valor
}

func main() {
	p := pessoa{nome: "Gui", saldo: 1000.}
	fmt.Println(p.getSaldo())
	fmt.Println(p.depositar(2000))
}
