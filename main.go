package main

import (
	"errors"
	"fmt"

	"golang-excercise/internal/entity"
	// "golang-excercise/internal/repository"
)

// ===== O Service define o que ele PRECISA (Contrato)
// ===== Ao realizar a implantação da interface precisa satisfazer o que foi estabelecido
type AlunoRepository interface {
    Save(aluno entity.Aluno) error
}
// =====

// ===== Implementação do repository
type MemoryAlunoRepository struct{}

func (r *MemoryAlunoRepository) Save(aluno entity.Aluno) error {
	fmt.Printf("Aluno %v salvo com sucesso!\n", aluno)
	return nil
}
// =====

// ===== Isso faz pate do arquivo useCase.go
// ===== Abstração usada no UseCase
type Service struct {
	// Abstração onde repo representa AlunoRepository
    repo AlunoRepository
}
// UseCase da aplicação, chama o repository
func (s *Service) Matricular(nome string) error {
    // Regra de negócio: Aluno não pode ter nome vazio
    if nome == "" {
        return errors.New("nome inválido")
    }
	fmt.Printf("O valor do nome é: %v\n", nome)

	// Conversão de tipo declarando nova variável
    aluno := entity.Aluno{Nome: nome}
	// aluno := "Marco"
	fmt.Printf("O valor de novoAluno é: %v, e o tipo é %T\n", aluno, aluno)
	
	err := s.repo.Save(aluno); if err != nil {
		fmt.Printf("Erro\n")
	}

	return nil
}
// =====

// ===== Construtir, bootstraper
func NewService(repo AlunoRepository) *Service {
	return &Service{
		repo: repo,
	}
}
// =====

// ===== 
func main() {

	repo := &MemoryAlunoRepository{}

	service := NewService(repo)

	err := service.Matricular("Marco")
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
// ===== 
