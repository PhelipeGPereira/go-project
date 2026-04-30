package main

import "fmt"

var usuario []Usuario
var proximoID = 1

func criar(nome string, email string, idade int) Usuario {
	u := Usuario{
		ID:    proximoID,
		Nome:  nome,
		Email: email,
		Idade: idade,
	}
	proximoID++
	usuario = append(usuario, u)
	return u
}

func listar() []Usuario {
	return usuario
}

func buscarPorID(id int) (Usuario, error) {
	for _, u := range usuario {
		if u.ID == id {
			return u, nil
		}
	}
	return Usuario{}, fmt.Errorf("usuário %d não encontrado", id)
}

func deletar(id int) error {
	for i, u := range usuario {
		if u.ID == id {
			usuario = append(usuario[:i], usuario[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("usuário %d não encontrado", id)
}

func atualizar(id int, novoNome string, novoEmail string) error {
	for i := range usuario {
		if usuario[i].ID == id {
			usuario[i].Nome = novoNome
			usuario[i].Email = novoEmail
			return nil
		}
	}
	return fmt.Errorf("Usuário %d não encontrado", id)
}
