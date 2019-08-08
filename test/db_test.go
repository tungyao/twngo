package test

import (
	"../tndb"
	"testing"
)

var db tndb.FUNC = new(tndb.DB).Use("tung_db")

func TestDb(t *testing.T) {
	//db.Insert("t_user").Key([]string{`NULL`, "123123", `NULL`}).Done()
	//db.Update("t_user").Key(map[string]string{"email":"asdasdas"}).Where(map[string]string{"id":"10"}).Done()
	db.Select("t_user").FindOne("id", "email")
}
