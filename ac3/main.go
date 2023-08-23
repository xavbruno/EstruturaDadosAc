package main

import (
	"fmt"
	"projetoAc3/utils"
	"bufio"
	"os"
	"strings"
)

func main() {
	var contatos [5]*utils.Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover, (3) para editar e-mail, (4) para exibir contatos ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')
			nome = strings.TrimSpace(nome)

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			utils.AdicionaContato(&utils.Contato{Nome: nome, Email: email}, &contatos)
		case "2":
			utils.ExcluiContato(&contatos)
		case "3":
			fmt.Println("Contatos:")
			for i, contato := range contatos {
				if contato != nil {
					fmt.Printf("[%d] %s - %s\n", i, contato.Nome, contato.Email)
				}
			}

			fmt.Print("Informe o índice do contato a ter o e-mail editado: ")
			var indice int
			fmt.Scanln(&indice)

			if indice >= 0 && indice < len(contatos) && contatos[indice] != nil {
				fmt.Print("Informe o novo email: ")
				fmt.Scanln(&email)

				contatos[indice].AlterarEmail(email)
				fmt.Println("E-mail editado com sucesso!")
			} else {
				fmt.Println("Índice inválido ou contato inexistente.")
			}
		case "4":
			fmt.Println("Contatos:")
			for i, contato := range contatos {
				if contato != nil {
					fmt.Printf("[%d] %s - %s\n", i, contato.Nome, contato.Email)
				}
			}
		default:
			return
		}

		fmt.Println(contatos)
	}
}
