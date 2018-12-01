package main

import (
	"fmt"
)

func main() {

	fmt.Println("Digite o Primeiro Numero:")
	var numero1 int64
	_, err := fmt.Scanln(&numero1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Primeiro Numero é: ,", numero1)

	fmt.Println("Digite o Segundo Numero:")
	var numero2 int64
	_, err1 := fmt.Scanln(&numero2)
	if err1 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Segundo Numero é:", numero2)

	fmt.Println("Digite a operação:")
	var operacao string
	_, err3 := fmt.Scanln(&operacao)
	if err3 != nil {
		fmt.Println(err)
		return
	}
	switch operacao {
	case "+":
		var total int64
		total = numero1 + numero2
		fmt.Println("O Valor Total é:", total)
	case "-":
		var total int64
		total = numero1 - numero2
		fmt.Println("O Valor Total é:", total)
	case "/":
		var total int64
		total = numero1 / numero2
		fmt.Println("O Valor Total é:", total)
	case "*":
		var total int64
		total = numero1 + numero2
		fmt.Println("O Valor Total é:", total)
	}

}
