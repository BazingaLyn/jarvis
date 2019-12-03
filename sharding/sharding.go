package sharding

import (
	"database/sql"
	"fmt"
)

type DB struct {
	*sql.DB
	driverName string
}

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{DB: db, driverName: driverName}, err
}

func (db *DB) ShardingExec(query string, args ...interface{}) (sql.Result, error) {
	fmt.Println(query)
	return MustExec(db, query, args...)
}

func MustExec(db *DB, query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}
