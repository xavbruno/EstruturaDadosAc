package main

import "fmt"

func inserirOrdenado(lista []int, valor int) []int {
	n := len(lista)
	novaLista := make([]int, n+1)
	inserido := false

	for i := 0; i < n+1; i++ {
		if !inserido && (i == n || valor < lista[i]) {
			novaLista[i] = valor
			inserido = true
		} else if inserido {
			novaLista[i] = lista[i-1]
		} else {
			novaLista[i] = lista[i]
		}
	}

	return novaLista[:n+1]
}

func main() {
	lista := []int{1, 3, 5, 7, 9}
	lista = inserirOrdenado(lista, 4)
	fmt.Println(lista)
}

