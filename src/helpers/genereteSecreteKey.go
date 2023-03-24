package helpers

import (
	"api/src/respostas"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

const POST = http.MethodPost
const GET = http.MethodGet
const PUT = http.MethodPut
const DELETE = http.MethodDelete

func Print(value interface{}) {
	if value != "" {
		fmt.Println(value)
	} else {
		fmt.Print("ERRO AO PRINTAR NA TELA POIS  O VALOR do parametro ESTA VAZIO")
	}
}

func GenereteNewSecreteKey() (string, error) {
	chave := make([]byte, 64)

	if _, err := rand.Read(chave); err != nil {
		return "", err
	}

	str64 := base64.StdEncoding.EncodeToString(chave)

	return str64, nil
}

func Err(w http.ResponseWriter, status int, err error) {
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
}
