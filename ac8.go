package main

import (
	"fmt"
)

type No struct {
	Chave  int
	Esquerda *No
	Direita *No
}

func novoNo(chave int) *No {
	return &No{Chave: chave, Esquerda: nil, Direita: nil}
}

func Inserir(raiz *No, chave int) *No {
	if raiz == nil {
		return novoNo(chave)
	}

	if chave < raiz.Chave {
		raiz.Esquerda = Inserir(raiz.Esquerda, chave)
	} else if chave > raiz.Chave {
		raiz.Direita = Inserir(raiz.Direita, chave)
	}

	return raiz
}

func Buscar(raiz *No, chave int) *No {
	if raiz == nil || raiz.Chave == chave {
		return raiz
	}

	if chave < raiz.Chave {
		return Buscar(raiz.Esquerda, chave)
	}

	return Buscar(raiz.Direita, chave)
}

func EmOrdem(raiz *No) {
	if raiz != nil {
		EmOrdem(raiz.Esquerda)
		fmt.Printf("%d ", raiz.Chave)
		EmOrdem(raiz.Direita)
	}
}

func main() {
	var raiz *No

	raiz = Inserir(raiz, 50)
	Inserir(raiz, 30)
	Inserir(raiz, 20)
	Inserir(raiz, 40)
	Inserir(raiz, 70)
	Inserir(raiz, 60)
	Inserir(raiz, 80)

	fmt.Println("Árvore em ordem:")
	EmOrdem(raiz)
	fmt.Println()

	chaveParaBuscar := 40
	resultado := Buscar(raiz, chaveParaBuscar)
	if resultado != nil {
		fmt.Printf("Chave %d encontrada na árvore.\n", chaveParaBuscar)
	} else {
		fmt.Printf("Chave %d não encontrada na árvore.\n", chaveParaBuscar)
	}
}
