package repro

import (
	"gopkg.in/mgo.v2/bson"
	"cursodego/bancomongo/model"

)

func ObtemLocal(codigotelefone int) (local model.Local, err error) {

	// Performance: Sempre copiar a Sessão e dar um defer close
	sessao := SessaoMongo.Copy()
	defer sessao.Close()
	colecao := sessao.DB("cursodego").C("local")
	err = colecao.Find(bson.M{"telcod":codigotelefone}).One(&local)
	return

}