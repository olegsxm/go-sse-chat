package db

import (
	"github.com/olegsxm/go-sse-chat/ent"
)

type Db struct {
	sql *ent.Client
}

func (d *Db) SQL() *ent.Client {
	return d.sql
}

func NewDB() *Db {
	return &Db{
		sql: newSqlClient(),
	}
}
