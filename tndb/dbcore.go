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
)

type FUNC interface {
	Use(string) *DB
	Insert(string) *DB
	Update(string) *DB
	Key(interface{}) *DB
	Where(map[string]string) *DB
	Done() bool
}

type DB struct {
	op    int
	db    string
	table string
	sql   string
	kel   *sql.DB
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
func (d *DB) Done() bool {
	stmt, err := d.kel.Prepare(d.sql)
	fmt.Println(err)
	res, _ := stmt.Exec()
	id, _ := res.LastInsertId()
	fmt.Println(id)
	return true
}
