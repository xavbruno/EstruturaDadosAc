package utils

import (
	"fmt"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func AdicionaContato(c *Contato, a *[5]*Contato) {
	for ind, contato := range *a {
		if contato == nil {
			(*a)[ind] = c
			break
		}
	}
}

func ExcluiContato(a *[5]*Contato) {
	if a[0] == nil {
		fmt.Println("Lista de contatos vazia!")
		return
	}

	for ind, contato := range *a {
		if contato == nil {
			(*a)[ind-1] = nil
		}
	}
}