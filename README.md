/meu-projeto
  /internal
    /handler      <-- (O "Roteador") Recebe o HTTP, valida JSON, responde 200/400/500.
    /service      <-- (O "Cérebro") Regras de negócio. Chama o Repository.
    /repository   <-- (O "Almoxarife") Consultas SQL (GORM, sqlx ou standard).
    /entity       <-- (O "Modelo") As structs puras (ex: type User struct).
  main.go         <-- Onde tudo se conecta (Injeção de Dependência).

  Handler Usa pacotes como net/http, Gin ou Fiber. Não colocar lógica de banco aqui. Apenas converter o Body para uma Struct.

  Service (Domain) É onde o Go brilha com Interfaces. Definir Interfaces. O Service não sabe qual banco existe, ele só sabe que existe um "salvador de dados".

  Repository Implementa a Interface do Service. Escrever o SQL ou usar um ORM. É a única camada que "conhece" a tecnologia do banco.





Exemplo Prático (Snippet Go)
​ "enxergar" a abstração, mostre como o Service (Domínio) se isola usando uma Interface:

// internal/entity/aluno.go
type Aluno struct {
    ID   int
    Nome string
}

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