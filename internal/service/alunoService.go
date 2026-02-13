package service

import (
	// "log"
	"errors"

	"golang-excercise/internal/entity"
)

// internal/service/aluno_service.go
// O Service define o que ele PRECISA (Contrato)
type AlunoRepository interface {
    Save(aluno entity.Aluno) error
}

type Service struct {
    repo AlunoRepository // Abstração!
}

func (s *Service) Matricular(nome string) error {
    // Regra de negócio: Aluno não pode ter nome vazio
    if nome == "" {
        return errors.New("nome inválido")
    }
    novoAluno := entity.Aluno{Nome: nome}
    return s.repo.Save(novoAluno)
}
