package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println("Starter API on PORT: ", config.Porta)

	r := router.Gerar()

	var PORT = 5000

	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), r))
}
