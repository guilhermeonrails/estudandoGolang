package main

import "time"

func main() {
	go abc()
	println("Primeiro programa com concorrencia!")
	time.Sleep(1 * time.Second)
}
func abc() {
	for i := byte('a'); i <= byte('z'); i++ {
		println(string(i))
	}
}
