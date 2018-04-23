package main

import (
	"fmt"
)

func main() {

	meses := 6
	situacao := true
	cidade := "Teste"

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo.")
	}

	if situacao {
		fmt.Println("Ele esta devendo")
	}

	if !situacao {
		fmt.Println("Ele esta adimplente")
	}

	if cidade == "Teste" {
		fmt.Println("Cidade é Teste")
	}

	if descricao, status := haQuantoTemoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente?", descricao)
		fmt.Println("Qual o status do cliente?", status)
		return
	}
	//a instrução abaixo não compila
	//fmt.Println("Obrigado por nos consultar!!!", descricao)
	fmt.Println("Obrigado por nos consultar!!!")
}

func haQuantoTemoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente esta devendo"
		return
	}
	descricao = "O cliente esta em dia"
	return
}
