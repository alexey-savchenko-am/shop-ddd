package persistence

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SqlxQueryDB struct {
	db *sqlx.DB
}

func NewSqlxQueryDB(db *sqlx.DB) *SqlxQueryDB {
	return &SqlxQueryDB{db: db}
}

func (q *SqlxQueryDB) formatQuery(query string, args ...any) (string, []any, error) {
	named, params, err := sqlx.Named(query, args)
	if err != nil {
		return "", nil, fmt.Errorf("build query: %w", err)
	}
	named = q.db.Rebind(named)
	return named, params, nil
}
func (q *SqlxQueryDB) Select(dest interface{}, query string, args ...any) error {
	named, params, err := q.formatQuery(query, args...)
	if err != nil {
		return err
	}
	return q.db.Select(dest, named, params...)
}

func (q *SqlxQueryDB) Get(dest interface{}, query string, args ...any) error {
	named, params, err := q.formatQuery(query, args...)
	if err != nil {
		return err
	}
	return q.db.Get(dest, named, params...)
}
