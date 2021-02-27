package repro

import (
	"gopkg.in/mgo.v2"
)

var SessaoMongo *mgo.Session


func AcreSessaoMongo() (err error) {

	err = nil
	SessaoMongo, err = mgo.Dial("mongobd://localhost:27017/cursodego")
	return

}