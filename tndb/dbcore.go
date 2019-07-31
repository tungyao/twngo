package tndb

import "fmt"

const (
	INSERT = iota
	UPDATE
	DELETE
)

type FUNC interface {
	Use(string) DB
	Insert(string) DB
	Update(string) DB
	Key(interface{}) DB
	Where(map[string]string) DB
	Done() bool
}

type DB struct {
	op    int
	db    string
	table string
	sql   string
}

//TODO 使用数据库 Use
func (d DB) Use(dbname string) DB {
	d.db = dbname
	return d
}

//TODO 插入数据 HEAD
func (d DB) Insert(table string) DB {
	d.op = INSERT
	d.table = table
	d.sql = "insert into " + table
	return d
}

//TODO 升级数据
func (d DB) Update(table string) DB {
	d.op = UPDATE
	d.table = table
	d.sql = "update " + table
	return d
}

//TODO 插入数据 / 升级 KEY
func (d DB) Key(k interface{}) DB {
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
func (d DB) Where(key map[string]string) DB {
	d.sql += " where " + ConvertMapString(key)
	return d
}

//TODO 数据
func (d DB) Done() bool {
	fmt.Println(d)
	return true
}