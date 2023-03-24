package models

type Seguidores struct {
	Usuario_ID  uint64 `json:"usuario_id, omitempty"`
	Seguidor_ID uint64 `json:"seguidor_id, omitempty"`
}
