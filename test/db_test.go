package test

import (
	"../tndb"
	"fmt"

	"testing"
)

var db tndb.FUNC = new(tndb.DB).Use("tung_db", "123456")

func TestDb(t *testing.T) {
	//db.Insert("user").Key([]string{`NULL`, "123123123", `NULL`}).Done()
	//db.Update("t_user").Key(map[string]string{"email":"asdasdas"}).Where(map[string]string{"id":"10"}).Done()
	data := db.Select("user").All("id", "email")
	fmt.Println(tndb.B2S(data))
}
