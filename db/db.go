package db

import (
	"database/sql"
	"net/url"

	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

type DB struct {
	db    *sql.DB
	dbmap *gorp.DbMap
}

func NewDB(uri string) (*DB, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(u.Scheme, uri)
	if err != nil {
		return nil, err
	}

	var dialect gorp.Dialect
	switch u.Scheme {
	case "postgres":
		dialect = gorp.PostgresDialect{}
	default:
		dialect = gorp.SqliteDialect{}
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: dialect}

	return &DB{
		dbmap: dbmap,
		db:    db,
	}, nil
}

func (db *DB) AddTableWithName(v interface{}, name string) interface {
	SetKeys(isAutoIncr bool, fieldNames ...string) *gorp.TableMap
} {
	return db.dbmap.AddTableWithName(v, name)
}

func (db *DB) Insert(v interface{}) error {
	return db.dbmap.Insert(v)
}

func (db *DB) SelectOne(v interface{}, query string, args ...interface{}) error {
	return db.dbmap.SelectOne(v, query, args...)
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.dbmap.Exec(query, args...)
}

func (db *DB) Close() error {
	return db.db.Close()
}