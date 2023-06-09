package models

import (
	"errors"
	"strings"
)

type Transactions struct {
	Id_cp            uint64  `json:"id_cp"`
	Name_cp          string  `json:"name_cp"`
	Amount_cp        float64 `json:"amount_cp"`
	Description_cp   string  `json:"description_cp"`
	Category_id      string  `json:"category_id"`
	Type_cp          string  `json:"type_cp"`
	PayStatus_cp     string  `json:"payStatus_cp"`
	DueDate_cp       string  `json:"dueDate_cp"`
	Qnt              uint64  `json:"qnt"`
	FirstInstallment string  `json:"firstInstallment"`
	Status_cp        uint64  `json:"status_cp"`
	Recorrente_cp    string  `json:"recorrente_cp"`
	Created_at       string  `json:"created_at"`
	Updated_at       string  `json:"updated_at"`
	Empresa_id       uint64  `json:"empresa_id"`
	User_id          uint64  `json:"user_id"`
	ClientID_cp      uint64  `json:"clientID_cp"`
}

func (t *Transactions) Check() error {
	err := t.validar()
	if err != nil {
		return err
	}

	t.formatar()
	return nil
}

func (t *Transactions) formatar() {
	t.Name_cp = strings.TrimSpace(t.Name_cp)
}

func (t *Transactions) validar() error {

	if len(t.Name_cp) > 100 {
		return errors.New("o nome da transação deve ter no máximo 100 caracteres")
	}

	if len(t.Name_cp) <= 3 {
		return errors.New("o nome da transação deve ter no minimo 3 ou mais caracteres")
	}

	return nil

}
