package middlewares

import (
	"net/http"

	pkgEntity "github.com/matheuslimab/artikoo/api/pkg/entity"
	"github.com/matheuslimab/artikoo/api/src/auth"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("-----------------------------------------------")
		// log.Printf("\n %s  %s  %s ", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

/*
* Autenticar verifica se o usuario esta fazendo requisicao autenticado
 */
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidarToken(r); err != nil {
			pkgEntity.Erro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
