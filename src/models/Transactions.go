package models

type Transactions struct {
	Id_cp            uint64 `json:"id_cp"`
	Name_cp          string `json:"name_cp"`
	Amount_cp        string `json:"amount_cp"`
	Description_cp   string `json:"description_cp"`
	Category_id      string `json:"category_id"`
	Type_cp          string `json:"type_cp"`
	PayStatus_cp     string `json:"payStatus_cp"`
	DueDate_cp       string `json:"due_date_cp"`
	Qnt              uint64 `json:"qnt"`
	FirstInstallment bool   `json:"first_installment"`
	Status_cp        uint64 `json:"status_cp"`
	Recorrente_cp    string `json:"recorrente_cp"`
	Created_at       string `json:"created_at"`
	Updated_at       string `json:"updated_at"`
	Empresa_id       uint64 `json:"empresa_id"`
	User_id          uint64 `json:"user_id"`
}
