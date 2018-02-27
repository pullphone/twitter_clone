package infrastructure

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type DB interface {
	Execute(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sqlx.Rows, error)
	Close()
}

var dbs = map[string]*db{}

type db struct {
	name string
	sqlxDb *sqlx.DB
}

func GetDatabase(name string) *db {
	if dbs[name] == nil {
		dbs[name] = &db{
			name: name,
			sqlxDb: dbOpen(name),
		}
	}

	return dbs[name]
}

func dbOpen(name string) (*sqlx.DB) {
	cfg := mysql.Config{}
	cfg.Net = "tcp"
	cfg.DBName = "twitter_clone"
	cfg.User = "root"
	//cfg.Passwd = ""
	cfg.ParseTime = true
	cfg.Loc, _ = time.LoadLocation("Asia/Tokyo")

	dsn := cfg.FormatDSN()
	db := sqlx.MustOpen("mysql", dsn)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(200)
	return db
}

func (db *db) Execute(statement string, args ...interface{}) (result sql.Result, err error) {
	result, err = db.sqlxDb.Exec(statement, args...)
	return
}

func (db *db) Query(statement string, args ...interface{}) (rows *sqlx.Rows, err error) {
	rows, err = db.sqlxDb.Queryx(statement, args...)
	return
}

func (db *db) Close() {
	name := db.name
	if dbs[name] != nil {
		dbs[name].Close()
		dbs[name] = nil
		db.sqlxDb = nil
	}
}
