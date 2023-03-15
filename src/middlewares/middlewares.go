package middlewares

import (
	"api/src/auth"
	"api/src/respostas"
	"fmt"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("-----------------------------------------------")
		log.Printf("\n %s  %s  %s ", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

/*
* Autenticar verifica se o usuario esta fazendo requisicao autenticado
 */
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidarToken(r); err != nil {
			respostas.Erro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
