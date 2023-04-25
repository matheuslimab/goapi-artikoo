package controllers

import (
	"api/src/models"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	var transactions models.Transactions
	if err := json.Unmarshal(body, &transactions); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println("+reactNative")
	fmt.Println(transactions)

	// db, err := database.Connect()
	// if err != nil {
	// 	respostas.Erro(w, http.StatusInternalServerError, err)
	// 	return
	// }
	// defer db.Close()

}
