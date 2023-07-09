package database

import (
	"database/sql"
	"fmt"

	"github.com/matheuslimab/artikoo/api/config"

	_ "github.com/go-sql-driver/mysql"
)

var stringConnec string

func Connect() (*sql.DB, error) {

	if config.ProductionServer {
		// conexao remota
		stringConnec = config.StringRemoteDb
	} else {
		// conexao local
		stringConnec = config.StringConexaoBanco
	}

	db, err := sql.Open("mysql", stringConnec)
	if err != nil {
		return nil, err
	}

	if config.ProductionServer {
		fmt.Println("A conexao ao banco de dados remota esta pronta!")
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
