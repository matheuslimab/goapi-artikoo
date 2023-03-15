package security

import "golang.org/x/crypto/bcrypt"

func Hash(senha string) ([]byte, error) {
	// defaultCost = 10
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func Verify(hashedSenha string, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSenha), []byte(senha))
}
