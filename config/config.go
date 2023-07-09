package config

import (
	"fmt"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	// chave usada para assinar o token de autenticação
	SecretKey []byte
)

func LoadConfigs() {

	defer fmt.Println("Iniciando configuração externa...")

	//var err error

	// if err = godotenv.Load(); err != nil {
	// 	fmt.Println(".env não foi encontrado!")
	// 	log.Fatal(err)
	// }

	// Porta, err = strconv.Atoi(os.Getenv("PORT"))
	// if err != nil {
	// 	fmt.Println("## erro na porta por favor insira outra porta! ##")
	// 	log.Fatal(err)
	// }

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		Db_user,
		Db_password,
		DB_name,
	)

	SecretKey = []byte(Secret_key)

}
