package tndb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	INSERT = iota
	UPDATE
	DELETE
	Select
)

type FUNC interface {
	Use(string) *DB

	Select(string) *DB
	All() *DB
	FindOne(column ...interface{}) map[string]interface{}
	Count() *DB
	IsExits() *DB

	Insert(string) *DB
	Update(string) *DB
	Key(interface{}) *DB
	Where(map[string]string) *DB
	Done() int64
}

type DB struct {
	op        int
	db        string
	table     string
	sql       string
	kel       *sql.DB
	formatSql map[string]string
}

func (d *DB) Select(tablename string) *DB {
	d.op = Select
	d.formatSql = make(map[string]string)
	d.formatSql["select"] = "select "
	d.formatSql["from"] = " from " + tablename
	return d
}

func (d *DB) All() *DB {
	d.sql += "*"

	return d
}

type data struct {
	id    int
	email string
}

func (d *DB) FindOne(column ...interface{}) map[string]interface{} {
	var tos string = ""
	if len(column) != 0 {
		for _, v := range column {
			tos += string(v.(string)) + ","
		}
	} else {
		tos = " * "
	}

	d.formatSql["column"] = tos[:len(tos)-1]
	d.formatSql["limit"] = " limit 1"
	d.sql = d.formatSql["select"] + d.formatSql["column"] + d.formatSql["from"] + d.formatSql["where"] + d.formatSql["limit"]
	fmt.Println(d.sql)
	rows, err := d.kel.Query(d.sql)
	toError(err)
	columns, _ := rows.Columns()
	length := len(columns)
	data := make(map[string]interface{})
	for rows.Next() {
		value := make([]interface{}, length)
		columnPointers := make([]interface{}, length)
		for i := 0; i < length; i++ {
			columnPointers[i] = &value[i]
		}
		rows.Scan(columnPointers...)
		for i := 0; i < length; i++ {
			columnName := columns[i]
			columnValue := columnPointers[i].(*interface{})
			data[columnName] = *columnValue
		}
	}
	fmt.Println(data)
	return data
}

func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
func (d *DB) Count() *DB {
	return d
}
func (d *DB) IsExits() *DB {
	return d
}

//TODO 使用数据库 Use
func (d *DB) Use(dbname string, pwd string) *DB {
	d.db = dbname
	db, err := sql.Open("mysql", "root:"+pwd+"@tcp(localhost)/"+dbname+"?charset=utf8")
	fmt.Println(err)
	d.kel = db
	return d
}

//TODO 插入数据 HEAD
func (d *DB) Insert(table string) *DB {
	d.op = INSERT
	d.table = table
	d.sql = "insert into " + table
	return d
}

//TODO 升级数据
func (d *DB) Update(table string) *DB {
	d.op = UPDATE
	d.table = table
	d.sql = "update " + table
	return d
}

//TODO 插入数据 / 升级 KEY
func (d *DB) Key(k interface{}) *DB {
	sl := keyForInsertOrUpdate(k)
	switch d.op {
	case INSERT:
		d.sql += sl
	case UPDATE:
		d.sql += sl
	}
	return d
}

//TODO
func (d *DB) Where(key map[string]string) *DB {
	d.sql += " where " + ConvertMapString(key)
	d.formatSql["where"] = " where " + ConvertMapString(key)
	return d
}

//TODO 数据
func (d *DB) Done() int64 {

	stmt, _ := d.kel.Prepare(d.sql)
	res, err := stmt.Exec()
	toError(err)
	id, err := res.LastInsertId()
	toError(err)
	return id

}
func toError(err error) {
	return
}
