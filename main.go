package main

import (
	"errors"
	"fmt"

	"golang-excercise/internal/entity"
)

// O Service define o que ele PRECISA (Contrato)
// Ao realizar a implantação da interface precisa satisfazer o que foi estabelecido
type AlunoRepository interface {
    Save(aluno entity.Aluno) error
}

type Service struct {
	// Abstração onde repo representa AlunoRepository
    repo AlunoRepository
}

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

func (s *Service) Repo(aluno entity.Aluno) error {
	fmt.Printf("O aluno %v foi salvo com sucesso no banco de dados\n", aluno)

	return nil
}

// Crian instância *Service e retorna &Service
func NewApp() *Service {
	return &Service{}
}

func main() {
	// Declara a variável service do tipo *Service
	service := NewApp()

	fmt.Printf("O tipo do nome: %T\n", service)

	// Chama o método matricular do tipo *Service que retorna error ou nil
	// Usa a variável err para representar o error
	// Usar a variavel service para chamar Matricular que é do tipo *Service
	err := service.Matricular("Marco"); if err != nil {
		fmt.Printf("Deu erro no main.go ao chamar func Matricular")
	}
}
