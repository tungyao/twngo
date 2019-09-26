package test

import (
	"../tndb"
	"fmt"
	"testing"
)

var db tndb.FUNC = new(tndb.DB).Use("test", "123456")

func TestDb(t *testing.T) {
	//db.Insert("user").Key([]string{`NULL`, "123123123", `NULL`}).Done()
	//db.Update("t_user").Key(map[string]string{"email":"asdasdas"}).Where(map[string]string{"id":"10"}).Done()
	//data := db.Select("test").Where(map[string]string{"id":"2"}).FindOne("id", "name")
	//fmt.Println(data)
	//data2:=db.Select("test").All("id","name")
	//fmt.Println(data2)
	sql := "insert into test set name='hello'"
	la, rw := db.Command(sql).Execute()
	fmt.Println(la, rw)
}
