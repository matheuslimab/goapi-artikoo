package repository

import (
	"database/sql"

	"github.com/matheuslimab/artikoo/api/src/models"
)

type DB struct {
	db *sql.DB
}

func NewRepositoryTransactionsDB(db *sql.DB) *DB {
	return &DB{db}
}

func (repository DB) GetTransactions(transactionID uint64) (models.Transactions, error) {
	row, err := repository.db.Query(`SELECT id_cp, name_cp, amount_cp, description_cp, category_id, type_cp, payStatus_cp, dueDate_cp, qnt, firstInstallment, status_cp, recorrente_cp, empresa_id, user_id, clientID_cp FROM cp_cr WHERE id_cp ORDER BY id_cp DESC LIMIT 0,1`)
	if err != nil {
		return models.Transactions{}, err
	}
	defer row.Close()

	var transaction models.Transactions

	if row.Next() {
		err := row.Scan(
			&transaction.Id_cp,
			&transaction.Name_cp,
			&transaction.Amount_cp,
			&transaction.Description_cp,
			&transaction.Category_id,
			&transaction.Type_cp,
			&transaction.PayStatus_cp,
			&transaction.DueDate_cp,
			&transaction.Qnt,
			&transaction.FirstInstallment,
			&transaction.Status_cp,
			&transaction.Recorrente_cp,
			&transaction.Empresa_id,
			&transaction.User_id,
			&transaction.ClientID_cp,
		)
		if err != nil {
			return models.Transactions{}, err
		}
	}

	return transaction, nil
}

func (repository DB) CreateTransaction(transaction models.Transactions) (uint64, error) {
	statament, err := repository.db.Prepare("INSERT INTO cp_cr  (name_cp, amount_cp, description_cp, category_id, type_cp, payStatus_cp, dueDate_cp, installments_start_cp, installments_end_cp, qnt, firstInstallment, status_cp, recorrente_cp, empresa_id, user_id, clientID_cp) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statament.Close()

	result, err := statament.Exec(
		transaction.Name_cp,
		transaction.Amount_cp,
		transaction.Description_cp,
		transaction.Category_id,
		transaction.Type_cp,
		transaction.PayStatus_cp,
		transaction.DueDate_cp,
		"0",
		"0",
		transaction.Qnt,
		transaction.FirstInstallment,
		transaction.Status_cp,
		transaction.Recorrente_cp,
		transaction.Empresa_id,
		transaction.User_id,
		transaction.ClientID_cp,
	)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}
