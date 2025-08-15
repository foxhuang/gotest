package service

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

type Service struct {
	DB *sql.DB
}

func NewService() (*Service, error) {
	connString := "sqlserver://sa:Fox13141516@@mssql:1433?database=hylibTEST"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Service{DB: db}, nil
}
