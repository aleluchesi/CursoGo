package model

import (
	//"database/sql"
)

type Local struct {
	Pais		string			`jason:"pais" db:"country" bson:"contry"`
	Cidade		string 	`jsaon:"cidade" db:"city" bson:"city"`		
	Codigotel 	int				`json:"cod_tel" db:"telcod" bson:"telcod"`
}