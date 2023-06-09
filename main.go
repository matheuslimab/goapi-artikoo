package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matheuslimab/artikoo/api/config"
	"github.com/matheuslimab/artikoo/api/router"
)

func main() {
	config.LoadConfigs()
	fmt.Println("Starter API on PORT: ", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
