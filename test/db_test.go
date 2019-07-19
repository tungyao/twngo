package test

import (
	"../twngo/dbcore"
	"testing"
)



func TestDb(t *testing.T) {
	key := map[string]interface{}{"id":2,"email":"2250169694@qq.com"}
	dbcore.Use("tung_db").Insert("t_user").Key(key).Save()
}