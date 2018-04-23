package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3
	fmt.Print("O número ", numero, " se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	case 4:
		fmt.Println("quatro.")
	}

	fmt.Println("Você é da família do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough //continue
	case "freebsd":
		fallthrough //continue
	case "linux":
		fmt.Println("Sim!!")
	default:
		fmt.Println("Não sei")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("OK, você pode dormir até mais tarde!")
	default:
		fmt.Println("Não pode dormir até mais tarde")
	}

	numero = 9
	fmt.Println("Este número cabe num dígito?")
	switch {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve dois digitos...")
	case numero >= 100:
		fmt.Println("Número é muito grande")
	}
}
