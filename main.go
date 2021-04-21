package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/soucer10/bank/modelos"
)

func main() {

	t1:=modelos.Transação{2,1,35}
	t2:=modelos.Transação{2,3,35}

	println(t1.RealizarTrasacao())
	println(t2.RealizarTrasacao())

}