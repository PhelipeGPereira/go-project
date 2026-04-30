package main

import (
	"errors"
	"strings"
)

type Usuarios struct {
	ID    int
	Nome  string
	Email string
	Idade int
}

func validar(u Usuarios) error {
	if strings.TrimSpace(u.Nome) == "" {
		return errors.New("nome é obrigatório")
	}
	at := strings.Index(u.Email, "@")
	if at < 1 || !strings.Contains(u.Email[at:], ".") {
		return errors.New("email inválido")
	}
	if u.Idade < 0 || u.Idade > 120 {
		return errors.New("idade deve estar entre 0 e 120")
	}
	return nil
}
