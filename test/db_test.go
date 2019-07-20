package test

import (
	"../twngo/tgdb"
	"testing"
)

var db tgdb.FUNC = new(tgdb.DB).Use("tung_db")

func TestDb(t *testing.T) {
	key := map[string]string{"id": "2", "email": "2250169694@qq.com"}

	db.Insert("t_user").Key(key).Done()
	db.Insert("t_user").Key([]string{"NULL", "test@test.com", "NULL"}).Done()
	db.Update("t_user").Key(key).Where(map[string]string{"id": "1"}).Done()
}
func TestConvertString(t *testing.T) {
	//key := map[string]string{"id": "2", "email": "2250169694@qq.com"}
	keys := []string{"NULL", "test@test.com", "NULL"}
	//s :=tgdb.ConvertMapString(key)
	ss := tgdb.ConvertArrayString(keys)
	t.Log(ss)
}
