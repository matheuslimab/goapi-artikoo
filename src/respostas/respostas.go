package respostas

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// JSON retorna uma resposta em JSON para a requisicao
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		fmt.Println(dados)
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			fmt.Println("dentro do if de err em NewEncoder")
			log.Fatal(erro)
		}
	}
}

// erro retorna um erro em formato json
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		erro.Error(),
	})
}
