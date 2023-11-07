package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values ($1, $2, $3) returning id",
	)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	rows, erro := statement.Query(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}
	defer rows.Close()

	var IdPublicacao uint64
	if rows.Next() {
		if erro := rows.Scan(&IdPublicacao); erro != nil {
			return 0, erro
		}
	}

	return IdPublicacao, nil
}
