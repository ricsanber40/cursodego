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

	cidades := make([]string, 5)
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Marrocos"
	cidades[3] = "Tokio"
	cidades[4] = "Singapura"

	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] é %s\r\n", indice, cidade)
	}

	capitaisAsia := cidades[3:5]
	fmt.Println(capitaisAsia)

	items1 := cidades[:2]
	fmt.Println(items1)

	items2 := cidades[len(cidades)-2:]
	fmt.Println(items2)

	indiceARetirar := 2
	items3 := cidades[:indiceARetirar]
	items3 = append(items3, cidades[indiceARetirar+1:]...)
	fmt.Println(items3)
	copy(cidades, items3)
	fmt.Println(cidades)

}
