package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
	"log"
)

var err error

type usuarios struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// essa funcao é responsavel por inserir um usuario no banco de dados
func (repository usuarios) Criar(usuario models.Usuario) (uint64, error) {

	statament, err := repository.db.Prepare("INSERT INTO usuarios(nome, nick, email, senha) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}

	defer statament.Close()

	result, err := statament.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	lastIdInsert, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return uint64(lastIdInsert), nil
}

func (repository usuarios) VerifyMail(email string) (models.Usuario, error) {

	fmt.Println("Debug: verifyMail")

	rows, err := repository.db.Query("SELECT email FROM usuarios WHERE email = ?", email)
	if err != nil {
		fmt.Println("error na linha 52 verifyMail")
		fmt.Println(err)
		return models.Usuario{}, err
	}
	defer rows.Close()

	var usuario models.Usuario

	for rows.Next() {
		err := rows.Scan(
			&usuario.Email,
		)
		if err != nil {
			fmt.Println(err)
			return models.Usuario{}, err
		}
	}

	return usuario, err
}

func (repository usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	rows, err := repository.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? or nick LIKE ?",
		nomeOuNick,
		nomeOuNick,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var usuarios []models.Usuario

	for rows.Next() {

		var usuario models.Usuario
		if erro := rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

func (repository usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	rows, err := repository.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", ID)
	if err != nil {
		return models.Usuario{}, err
	}
	defer rows.Close()

	var usuario models.Usuario

	if rows.Next() {
		if erro := rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repository usuarios) AtualizarUsuario(ID uint64, usuario models.Usuario) error {
	statament, err := repository.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statament.Close()

	if _, err = statament.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository usuarios) DeletarUsuario(ID uint64) error {
	statament, err := repository.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		return err
	}
	defer statament.Close()

	if _, err := statament.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository usuarios) BuscarPorEmail(email string) (models.Usuario, error) {

	rows, err := repository.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer rows.Close()

	var usuario models.Usuario
	if rows.Next() {
		if erro := rows.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}
