package test

import (
	"../tndb"
	"testing"
)

var db tndb.FUNC = new(tndb.DB).Use("tung_db")

func TestDb(t *testing.T) {
	db.Insert("t_user").Key([]string{`NULL`, "123123", `NULL`}).Done()

}
