package infrastructure

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var db *sqlx.DB

func DBInit() {
	cfg := mysql.Config{}
	cfg.Net = "tcp"
	cfg.DBName = "twitter_clone"
	cfg.User = "root"
	//cfg.Passwd = ""
	cfg.ParseTime = true
	cfg.Loc, _ = time.LoadLocation("Asia/Tokyo")

	dsn := cfg.FormatDSN()
	db = sqlx.MustOpen("mysql", dsn)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(200)
}

func DBClose() {
	if db != nil {
		db.Close()
	}
}

func DBExecute(statement string, args ...interface{}) (result sql.Result, err error) {
	result, err = db.Exec(statement, args...)
	return
}

func DBQuery(statement string, args ...interface{}) (rows *sqlx.Rows, err error) {
	rows, err = db.Queryx(statement, args...)
	return
}
