package db

import "database/sql"

type DB struct {
	sql *sql.DB
}

func (d *DB) Sql() *sql.DB {
	return d.sql
}

func New() DB {
	return DB{
		sql: newSqliteConnection(),
	}
}
