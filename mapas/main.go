package main

import (
	"fmt"
	"structs_avancado/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["casa amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["casa amarela"]
	if achou {
		fmt.Printf("Sim, achei:  %+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto azul"
	apto.X = 19
	apto.Y = 26
	apto.SetValor(50000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\r\n", chave, imovel)
	}

	_, achou = cache["casa amarela"]
	if achou {
		delete(cache, "casa amarela")
	}

	fmt.Println("Quantos itens há no cache? ", len(cache))

}
