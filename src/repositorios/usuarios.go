package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO USUARIOS (nome, nick, email, senha) VALUES($1, $2, $3, $4) RETURNING ID")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Query(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}
	defer result.Close()

	var insertedId int64
	if result.Next() {
		erro := result.Scan(&insertedId)
		if erro != nil {
			return 0, erro
		}
	}

	return uint64(insertedId), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	statement, erro := repositorio.db.Prepare(`SELECT id, nome, nick, email, criado_em FROM USUARIOS WHERE (NOME ILIKE $1)OR(NICK ILIKE $2)`)

	if erro != nil {
		return nil, erro
	}
	defer statement.Close()

	linhas, erro := statement.Query(nomeOuNick, nomeOuNick)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.Id,
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

func (repositorio Usuarios) BuscarUsuario(id int64) (*modelos.Usuario, error) {

	linha, erro := repositorio.db.Query("select id, nome, nick, email, criado_em from usuarios where id = $1", id)
	if erro != nil {
		return &modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return &modelos.Usuario{}, erro
		}
	} else {
		return &modelos.Usuario{}, nil
	}

	return &usuario, nil
}

func (repositorio Usuarios) AtualizarUsuario(id int64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome = $1, nick = $2, email = $3 where id = $4")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id)

	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) DeletarUsuario(id int64) error {
	_, erro := repositorio.db.Exec("delete from usuarios where id = $1", id)
	if erro != nil {
		return erro
	}
	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = $1", email)

	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Id, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("insert into seguidores(usuario_id, seguidor_id) values($1, $2) on conflict do nothing")

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}
