package tgdb

import "fmt"

//TODO 定义Model接口 {
// dbname
// tablename
// table ->key
// table ->save()
// }
type _DB struct {
	db    string
	table string
	sql   string
}
type _KEY struct {
	_   interface{}
	KEY interface{}
}

var OperationalStatement map[string]string = map[string]string{
	"select": ""}

//TODO 使用数据库 Use
func Use(dbname string) *_DB {
	_d := &_DB{}
	_d.db = dbname
	return _d
}

//TODO 插入数据 HEAD
func (d _DB) Insert(table string) _DB {
	d.table = table
	return d
}

//TODO 插入数据 KEY
func (d _DB) Key(key map[string]interface{}) _DB {
	return d
}

//TODO 数据
func (d _DB) Save() bool {
	fmt.Println(d)
	return true
}
