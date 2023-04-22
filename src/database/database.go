package database

import (
	"api/src/config"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var production = true
var stringConnec string

func Connect() (*sql.DB, error) {

	if production {
		// conexao remota
		fmt.Println("Usando conexão remota agora...")
		stringConnec = os.Getenv("STRING_REMOTE")
	} else {
		// conexao local
		stringConnec = config.StringConexaoBanco
	}

	fmt.Println("-----------------------------STRING CONEXAO-----------------------------------")
	fmt.Println(stringConnec)
	fmt.Println("-------------------------------------------------------------------------------")

	db, err := sql.Open("mysql", stringConnec)
	if err != nil {
		return nil, err
	}

	if production {
		fmt.Println("A conexao ao banco de dados remota esta pronta!")
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
