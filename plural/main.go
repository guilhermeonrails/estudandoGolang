package main

var (
	message = "Meu primeiro programa!"
)

func main() {
	println(message)
}

func init() {
	message = "Função init"
}
