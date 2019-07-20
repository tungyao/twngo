package test

import (
	"../twngo/tgdb"
	"testing"
)

func TestDb(t *testing.T) {
	key := map[string]interface{}{"id": 2, "email": "2250169694@qq.com"}
	tgdb.Use("tung_db").Insert("t_user").Key(key).Save()
}
