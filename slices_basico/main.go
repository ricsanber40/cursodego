package main

import (
	"fmt"
)

func main() {

	var nums []int
	fmt.Println(nums, len(nums), cap(nums))

	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}
	capitais = append(capitais, "Brasilia")
	capitais[1] = "Brasília"
	fmt.Println("Qual a capacidade deste slice?", cap(capitais))
	for indice, cidade := range capitais {
		fmt.Printf("Capital[%d] é %s\r\n", indice, cidade)
	}

	cidades := make([]string, 4)
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Tokio"
	cidades[3] = "Singapura"
	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] é %s\r\n", indice, cidade)
	}
}
