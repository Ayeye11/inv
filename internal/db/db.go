package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgSQL struct {
	dsn string
}

func NewPgSQL(dsn string) *PgSQL {
	return &PgSQL{dsn}
}

func (psql *PgSQL) Run() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(psql.dsn), &gorm.Config{})
}
