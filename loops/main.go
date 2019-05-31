package main

func main() {
	//teste 1
	for i := 0; i < 5; i++ {
		println(i)
	}

	//teste 2
	nomes := []string{"Guilherme", "Lima", "Leticia", "Akemi"}
	for i, nome := range nomes {
		println(i, nome)
	}

	//teste 3
	nomesTeste2 := []string{"Guilherme", "Lima", "Leticia", "Akemi"}
	for _, nome := range nomesTeste2 {
		println(nome)
	}
}
