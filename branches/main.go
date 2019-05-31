package main

func main() {
	//teste 1
	idade := 20
	if idade >= 18 {
		println("Maior de idade")
	} else {
		println("Menor de idade")
	}

	if anos := 12; anos < 15 {
		println("Jovem")
	}

	//switch
	categoria := 16
	switch categoria {
	case 15:
		println("Sub 15")
	case 16:
		println("Sub 16")
	case 17:
		println("Sub 17")
	}

	//switch 2
	selecao := 18
	switch {
	case selecao <= 15:
		println("Sub 15")
	case selecao <= 17:
		println("Sub 17")
	case selecao >= 18:
		println("Sub 19")
	}
}
