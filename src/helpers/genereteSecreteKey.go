package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

func GenereteNewSecreteKey() (string, error) {
	chave := make([]byte, 64)

	if _, err := rand.Read(chave); err != nil {
		return "", err
	}

	str64 := base64.StdEncoding.EncodeToString(chave)

	return str64, nil
}
