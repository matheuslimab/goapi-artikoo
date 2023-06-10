package repository

import (
	"database/sql"

	"github.com/matheuslimab/artikoo/api/src/models"
)

type Db struct {
	db *sql.DB
}

func NewRepositorySegments(db *sql.DB) *dB {
	return &dB{db}
}

func (repository *dB) GetAllSegments() ([]models.Segments, error) {
	row, err := repository.db.Query(`SELECT id, nome FROM Segmentos`)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var segments []models.Segments

	for row.Next() {

		var segment models.Segments

		err := row.Scan(
			&segment.Id,
			&segment.Nome,
		)
		if err != nil {
			return nil, err
		}

		segments = append(segments, segment)

	}

	return segments, nil
}
