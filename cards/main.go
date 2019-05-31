package main

func main() {
	cards := deck{"Quatro de Espada", newCard()}
	cards = append(cards, "As de Ouro")

	cards.print()
}

//No Go, é necessário dizer o tipo do retorno da função
func newCard() string {
	return "Cinco de Ouro "
}
