package dao

import (
	"database/sql"
	"time"
	"github.com/go-sql-driver/mysql"
	"github.com/pullphone/twitter_clone/repository"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() repository.SqlHandler {
	cfg := mysql.Config{}
	cfg.Net = "tcp"
	cfg.DBName = "twitter_clone"
	cfg.User = "root"
	//cfg.Passwd = ""
	cfg.ParseTime = true
	cfg.Loc, _ = time.LoadLocation("Asia/Tokyo")

	dsn := cfg.FormatDSN()
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (repository.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}

	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (repository.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}

	sqlRow := new(SqlRow)
	sqlRow.Rows = rows
	return sqlRow, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
