package main

import "fmt"

type Contato struct{
	Nome string
	Email string
}

func (c Contato) mudaEmail() {
	c.Email = "carlos@gmail.com"
}

func main() {
	contato := Contato{
		Nome:	"Bruno",
		Email:	"bruno@gmail.com",
	}

	contato.mudaEmail()
	fmt.Println(contato)

	fmt.Println("-------------------")

	var c Contato
	c = Contato{
		Nome:	"Bruno",
	}
	var contatos [5]Contato

	contatos[0] = c
	for i, contato := range contatos {
		if(contato == Contato{}) {
			fmt.Println("achou espaço vazio no índice", i)
		}
	}
	fmt.Println("-------------------")

}

func adiciona(c Contato, a [5]Contato) [5]Contato {
	a[1] = c
	return a
}

func excluir(a [5]Contato) [5]Contato {
	for i, contato := range a {
		if(contato == Contato{}) {
			a[i - 1] = Contato{}
			break
		}
	}
	return a
}

