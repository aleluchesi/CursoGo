package repro

import (
	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"

)

var Db *sqlx.DB

func AbreConexaoSQL() (err error) {

	err = nil
	Db, err = sqlx.Open("mysql", "srvuser@tcp(127.0.0.1:3306)/cursodego")
	if err != nil {
		return
	}
	err = Db.Ping()
	return
}