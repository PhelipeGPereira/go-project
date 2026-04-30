package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func lerLinha(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func lerInt(prompt string) int {
	for {
		texto := lerLinha(prompt)
		n, err := strconv.Atoi(texto)
		if err != nil {
			fmt.Println("Digite um número válido")
			continue
		}
		return n
	}
}

func main() {
	for {
		fmt.Println("\n1. Cadastrar")
		fmt.Println("2. Listar")
		fmt.Println("3. Buscar por ID")
		fmt.Println("4. Atualizar")
		fmt.Println("5. Deletar")
		fmt.Println("0. Sair")

		opcao := lerInt("Escolha: ")

		switch opcao {
		case 1:
			nome := lerLinha("Nome: ")
			email := lerLinha("Email: ")
			idade := lerInt("Idade: ")
			u, err := criar(nome, email, idade)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("Cadastro: [%d] %s\n", u.ID, u.Nome)
			}

		case 2:
			usuarios := listar()
			if len(usuarios) == 0 {
				fmt.Println("Nenhum usuário cadastrado")
			}
			for _, u := range usuarios {
				fmt.Printf("[%d] %s | %s | %d anos\n", u.ID, u.Nome, u.Email, u.Idade)
			}

		case 3:
			id := lerInt("ID: ")
			u, err := buscarPorID(id)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("[%d] %s | %s | %d anos\n", u.ID, u.Nome, u.Email, u.Idade)
			}

		case 4:
			id := lerInt("ID: ")
			novoNome := lerLinha("Novo nome: ")
			novoEmail := lerLinha("Novo email: ")
			err := atualizar(id, novoNome, novoEmail)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Atualizado com sucesso")
			}

		case 5:
			id := lerInt("ID: ")
			err := deletar(id)
			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Deletado com sucesso")
			}

		case 0:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida")
		}
	}
}
