package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pkgEntity "github.com/matheuslimab/artikoo/api/pkg/entity"
	"github.com/matheuslimab/artikoo/api/src/database"
	"github.com/matheuslimab/artikoo/api/src/helpers"
	"github.com/matheuslimab/artikoo/api/src/models"
	"github.com/matheuslimab/artikoo/api/src/repository"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	var transactions models.Transactions
	if err := json.Unmarshal(body, &transactions); err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
		return
	}

	var transaction_id uint64

	dateStrFormatted, err := helpers.FormatterDateToGo(transactions.DueDate_cp)
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
	}

	transactions.DueDate_cp = dateStrFormatted

	err = transactions.Check()
	if err != nil {
		pkgEntity.Erro(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		pkgEntity.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	new_repository := repository.NewRepositoryTransactionsDB(db)

	if transactions.Qnt > 1 {

		transaction_id, err := new_repository.CreateTransaction(transactions)
		if err != nil {
			pkgEntity.Erro(w, http.StatusBadRequest, err)
			return
		}

		for i := 1; i < int(transactions.Qnt); i++ {

			fmt.Printf("P: %v \n", i)

			installments, err := new_repository.GetTransactions(transaction_id)
			if err != nil {
				pkgEntity.Erro(w, http.StatusBadRequest, err)
				return
			}

			installments.DueDate_cp, err = helpers.AddOneMonth(installments.DueDate_cp)
			if err != nil {
				pkgEntity.Erro(w, http.StatusBadRequest, err)
				return
			}

			transaction_id, err = new_repository.CreateTransaction(installments)
			if err != nil {
				pkgEntity.Erro(w, http.StatusBadRequest, err)
				return
			}
		}

		pkgEntity.JSON(w, http.StatusCreated, 1)

	} else {
		transaction_id, err = new_repository.CreateTransaction(transactions)
		if err != nil {
			pkgEntity.Erro(w, http.StatusBadRequest, err)
			return
		}

		transactions.Id_cp = transaction_id
		pkgEntity.JSON(w, http.StatusCreated, transaction_id)
	}
}
