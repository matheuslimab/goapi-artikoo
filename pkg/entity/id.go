package entity

import "github.com/google/uuid"

func GenerateNewID() string {
	return uuid.New().String()
}

func IsValidID(id uuid.UUID) (uuid.UUID, error) {
	id, err := uuid.Parse(id.String())
	return id, err
}
