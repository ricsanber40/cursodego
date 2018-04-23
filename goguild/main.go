package main

import (
	"encoding/json"
	"fmt"
	"structs_avancado/model"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(100000000); err != nil {
		fmt.Println("[main] Houve um erro na atribuição de valor a casa: ", err.Error())

		if err == model.ErrValorDeImovelAlto {
			fmt.Println("Corretor, atenção valor do imóvel muito alto")
		}
		return
	}

	fmt.Printf("Casa é: %v+\r\n", casa)

	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve erro na geração do objeto JSON: ", err.Error())
		return
	}
	fmt.Println("A casa em JSON:", string(objJSON))

}
