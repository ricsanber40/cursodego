package main

import (
	"fmt"
)

//Imovel é um struct
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 50000}
	fmt.Printf("O apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     17,
		Nome:  "Chacara",
		X:     25,
		valor: 25369,
	}

	fmt.Printf("A chacara é: %+v\r\n", chacara)

	casa.Nome = "Casa teste"
	casa.valor = 20
	casa.X = 10
	casa.Y = 20
	fmt.Printf("A casa agora é: %+v\r\n", casa)
}
