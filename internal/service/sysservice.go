package service

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlserver"

	_ "github.com/denisenkom/go-mssqldb"
)

type Service struct {
	DB *gorm.DB
}

func NewService() (*Service, error) {
	connString := "sqlserver://sa:Fox13141516@@172.29.32.1:1433?database=hylibTEST"
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Service{DB: db}, nil
}
