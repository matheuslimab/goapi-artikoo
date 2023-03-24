package repository

import (
	"api/src/models"
	"database/sql"
)

type publicacoesDB struct {
	db *sql.DB
}

func NewRepositoryPublicacoes(db *sql.DB) *publicacoesDB {
	return &publicacoesDB{db}
}

func (repository publicacoesDB) Criar(publicacao models.Publicacao) (uint64, error) {
	statament, err := repository.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statament.Close()

	result, err := statament.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	lastIDInsert, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInsert), nil
}
