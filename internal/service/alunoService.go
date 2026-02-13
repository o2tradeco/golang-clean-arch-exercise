/*
Service (Domain) É onde o Go brilha com Interfaces.
Definir Interfaces.
O Service não sabe qual banco existe, ele só sabe que existe um "salvador de dados".
*/
package service

import (

)

type AlunoRepository interface {
    Save(aluno entity.Aluno) error
}

type Service struct {
    repo AlunoRepository // Abstração!
}