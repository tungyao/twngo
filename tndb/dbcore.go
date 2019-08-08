package tndb

import (
	"database/sql"
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
	FindOne(column ...interface{}) *DB
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
	formatSql [4]string
}

func (d *DB) Select(tablename string) *DB {
	d.op = Select
	d.formatSql[0] = "select"
	d.formatSql[2] = "from " + tablename

	return d
}

func (d *DB) All() *DB {
	d.formatSql[1] = " * "

	return d
}

type data struct {
	id    int
	email string
}

func (d *DB) FindOne(column ...interface{}) *DB {
	var tos string
	for _, v := range column {
		tos += string(v.(string)) + ","

	}
	d.formatSql[1] = " " + tos[:len(tos)-1] + " "
	d.formatSql[3] = " limit 1"
	for _, v := range d.formatSql {
		d.sql += v
	}

	rows, err := d.kel.Query(d.sql)
	toError(err)
	columns, _ := rows.Columns()
	length := len(columns)
	for rows.Next() {
		value := make([]interface{}, length)
		columnPointers := make([]interface{}, length)
		for i := 0; i < length; i++ {
			columnPointers[i] = &value[i]
		}
		rows.Scan(columnPointers...)
		data := make(map[string]interface{})
		for i := 0; i < length; i++ {
			columnName := columns[i]
			columnValue := columnPointers[i].(*interface{})
			data[columnName] = *columnValue
		}
	}
	return d
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
func (d *DB) Use(dbname string) *DB {
	d.db = dbname
	db, _ := sql.Open("mysql", "root:123456@tcp(localhost)/tung_db?charset=utf8")
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
