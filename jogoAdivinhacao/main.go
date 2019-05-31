package main

import "fmt"

func main() {
	fmt.Println("**************************************")
	fmt.Println("*  Bem vindo ao jogo da adivinhação  *")
	fmt.Println("**************************************")

	var numeroSecreto = 42
	const numeroTentativas = 3
	//fmt.Println("Não conte para ninguém que o número secreto é", numeroSecreto)
	for index := 1; index <= numeroTentativas; index++ {
		fmt.Println("\nTentativa", index, "de", numeroTentativas)
		fmt.Println("Digite um número entre 0 e 50")
		var chute int
		fmt.Scanf("%d", &chute)

		if chute < 0 {
			fmt.Println("Você não pode digitar um número negativo")
			index--
			continue
		}
		/*
			Print imprimirá variáveis ​​numéricas e não incluirá uma quebra de linha no final.
			Printf não imprimirá variáveis ​​numéricas e não incluirá uma quebra de linha no final.
			Println imprimirá variáveis ​​numéricas e incluirá uma quebra de linha no final.
		*/
		var acertou = chute == numeroSecreto
		var chuteMaior = chute > numeroSecreto

		if acertou {
			fmt.Println("Você acertou!")
			fmt.Println("O número secreto era", numeroSecreto, "- Parabéns!")
			fmt.Println("**************************************")
			//for, para a repetição
			break
		} else if chuteMaior {
			fmt.Println("\nVocê errou, mas não desanime!")
			fmt.Println("Seu chute foi maior que o número secreto.")
		} else {
			fmt.Println("Seu chute foi menor que o número secreto.")
		}
		fmt.Println("**************************************")
	}
	fmt.Println("\nFim do jogo")
}
