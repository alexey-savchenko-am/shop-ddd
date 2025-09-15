package persistence

import "github.com/jmoiron/sqlx"

type SqlxQueryDB struct {
	db *sqlx.DB
}

func NewSqlxQueryDB(db *sqlx.DB) *SqlxQueryDB {
	return &SqlxQueryDB{db: db}
}

func (q *SqlxQueryDB) Select(dest interface{}, query string, args ...any) error {
	return q.db.Select(dest, query, args...)
}

func (q *SqlxQueryDB) Get(dest interface{}, query string, args ...any) error {
	return q.db.Get(dest, query, args...)
}
