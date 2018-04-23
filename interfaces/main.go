package main

import (
	"fmt"
	"projetos/interfaces/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	queroOuvirCacarejo(jojo)
	queroOuvirQuack(jojo)
}

func queroOuvirCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirQuack(g model.Pata) {
	fmt.Println(g.Quack())
}
