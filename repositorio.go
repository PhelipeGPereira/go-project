package main

import "fmt"

var usuario []Usuarios
var proximoID = 1

func criar(nome, email string, idade int) (Usuarios, error) {
	u := Usuarios{
		ID:    proximoID,
		Nome:  nome,
		Email: email,
		Idade: idade,
	}
	if err := validar(u); err != nil {
		return Usuarios{}, err
	}
	proximoID++
	usuario = append(usuario, u)
	return u, nil
}

func listar() []Usuarios {
	return usuario
}

func buscarPorID(id int) (Usuarios, error) {
	for _, u := range usuario {
		if u.ID == id {
			return u, nil
		}
	}
	return Usuarios{}, fmt.Errorf("usuário %d não encontrado", id)
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

func atualizar(id int, novoNome, novoEmail string) error {
	for i := range usuario {
		if usuario[i].ID == id {
			atualizado := Usuarios{ID: id, Nome: novoNome, Email: novoEmail, Idade: usuario[i].Idade}
			if err := validar(atualizado); err != nil {
				return err
			}
			usuario[i].Nome = novoNome
			usuario[i].Email = novoEmail
			return nil
		}
	}
	return fmt.Errorf("Usuário %d não encontrado", id)
}
