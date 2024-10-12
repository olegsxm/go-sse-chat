package storage

type Storage interface {
	Sql() Sql
}

type ST struct {
	sql Sql
}

func (s ST) Sql() Sql {
	return s.sql
}

func New() Storage {
	return ST{}
}
