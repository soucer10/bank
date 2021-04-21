package modelos

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Transação struct {

	Origem  uint
	Destino uint
	Valor 	uint64

}


func (t *Transação) RealizarTrasacao() error{

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err!=nil{
		fmt.Println(err.Error())
	}

	var c1 Conta
	var c2 Conta

	db.Table("conta").First(&c1,t.Origem)

	if c1.Saldo<t.Valor{
		return errors.New("Falta de saldo")
	}

	db.Table("conta").First(&c2,t.Destino)

	c1.Saldo=c1.Saldo-t.Valor
	c2.Saldo=c2.Saldo+t.Valor

	db.Table("conta").Where("id=?",t.Origem).Updates(c1)
	db.Table("conta").Where("id=?",t.Destino).Updates(c2)

	return nil
}