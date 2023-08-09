//Primos
package main

import "fmt"

func main() {
	fmt.Println(ePrimo(11))

	fmt.Println(diaSemana(3))

	fmt.Println(eBissexto(2012))
}

func ePrimo(x int) bool {
	var e_primo bool
	e_primo = true
	for i := 2; i < x; i++ {
		if x % i == 0 {
			e_primo = false
			fmt.Println("Não é primo")
		}
	}
	return e_primo
}

func diaSemana(dia int) string {

	switch dia{
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Dia Inválido"
	}
}

func eBissexto(x int) bool {
	var e_primo bool
	e_primo = true
	for i := 1; i <= x; i++ {
		if x % 4 > 0 {
			e_primo = false
		} else if x % 100 == 0 {
			e_primo = false
		}
	}
	return e_primo

}
