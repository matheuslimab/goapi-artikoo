package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	// chave usada para assinar o token de autenticação
	SecretKey []byte
)

func Carregar() {

	defer fmt.Println("Iniciando configuração externa...")

	var err error

	if err = godotenv.Load(); err != nil {
		fmt.Println(".env não foi encontrado!")
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		fmt.Println("## erro na porta por favor insira outra porta! ##")
		log.Fatal(err)
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
