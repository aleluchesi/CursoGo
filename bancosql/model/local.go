package model

import (
	"database/sql"
)

type Local struct {
	Pais		string			`jason:"pais" db:"country"`
	Cidade		sql.NullString 	`jsaon:"cidade" db:"city"`		
	Codigotel 	int				`json:"cod_tel" db:"telcod"`
}