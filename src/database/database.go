package database

import (
	"api/src/config"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var production = false
var stringConnec string

func Connect() (*sql.DB, error) {

	if production {
		// conexao remota
		stringConnec = os.Getenv("STRING_REMOTE")
	} else {
		// conexao local
		stringConnec = config.StringConexaoBanco
	}

	db, err := sql.Open("mysql", stringConnec)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
