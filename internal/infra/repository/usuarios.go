package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/matheuslimab/artikoo/api/src/models"
)

var err error

type usuarios struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// essa funcao é responsavel por inserir um usuario no banco de dados
func (repository usuarios) Criar(usuario models.Usuario) (string, error) {

	statament, err := repository.db.Prepare("INSERT INTO Users(id, nome, nick, email, senha) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	defer statament.Close()

	_, err = statament.Exec(usuario.ID, usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return usuario.ID, nil
}

func (repository usuarios) VerifyMail(email string) (models.Usuario, error) {
	rows, err := repository.db.Query("SELECT email FROM Users WHERE email = ?", email)
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

	rows, err := repository.db.Query("SELECT id, nome, nick, email, criadoEm FROM Users WHERE nome LIKE ? or nick LIKE ?",
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

func (repository usuarios) BuscarPorID(ID string) (models.Usuario, error) {
	rows, err := repository.db.Query("SELECT id, nome, nick, email, criadoEm FROM Users WHERE id = ?", ID)
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

func (repository usuarios) AtualizarUsuario(ID string, usuario models.Usuario) error {
	statament, err := repository.db.Prepare("UPDATE Users SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statament.Close()

	if _, err = statament.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository usuarios) DeletarUsuario(ID string) error {
	statament, err := repository.db.Prepare("DELETE FROM Users WHERE id = ?")
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

	rows, err := repository.db.Query("SELECT id, senha FROM Users WHERE email = ?", email)
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

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuário
func (repositorio usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
