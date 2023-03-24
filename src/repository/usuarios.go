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

// verificar se já estao se seguindo
func (repository usuarios) VerifyFollowing(usuarioID, seguidorID uint64) (models.Seguidores, error) {
	rows, err := repository.db.Query("SELECT usuario_id, seguidor_id FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?", usuarioID, seguidorID)
	if err != nil {
		return models.Seguidores{}, err
	}
	defer rows.Close()

	var seguidores models.Seguidores
	if rows.Next() {
		if err := rows.Scan(
			&seguidores.Usuario_ID,
			&seguidores.Seguidor_ID,
		); err != nil {
			return models.Seguidores{}, err
		}
	}

	return seguidores, nil
}

// Esse metodo registra no banco de dados que um usuario X seguira o usuario Y
func (repository usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statament, err := repository.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer statament.Close()

	if _, err := statament.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (repository usuarios) UnFollow(usuarioID, seguidorID uint64) error {
	statament, err := repository.db.Prepare("DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?")
	if err != nil {
		return err
	}
	defer statament.Close()

	if _, err := statament.Exec(usuarioID, seguidorID); err != nil {
		return err
	}

	return nil
}

func (repository usuarios) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repository.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
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

func (repository usuarios) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repository.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
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
