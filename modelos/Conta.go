package modelos

import (
	"gorm.io/gorm"
)

type Conta struct {
	gorm.Model
	Numerodaconta string
	Nome   		  string
	Saldo 		  uint64

}

//CreateConta Criação de Conta
func CreateConta(numerodaconta,nome string ) *Conta {

	return &Conta{Numerodaconta: numerodaconta,
		Nome: nome}

}

