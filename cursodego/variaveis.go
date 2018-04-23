package main

import "fmt"

var (
	//Endereco é um valor importante e tende a ser publico
	Endereco string
	//Telefone é publico
	Telefone            string  //Variavel Endereco e publico, telefone é privada
	quantidade, estoque int     //padrão 0
	comprou             bool    // padrão false
	valor               float64 // padrão 0.00
	palavras            rune
)

func main() {
	teste := "Valor de teste"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("teste: %s\r\n", teste)
}
