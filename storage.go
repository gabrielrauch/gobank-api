package main

import (
	"database/sql"
)

type Storage interface {
	CreateAccount(account *Account) error
	GetAccountByID(id int) (*Account, error)
	UpdateAccount(account *Account) error
	DeleteAccount(id int) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connectionString := "user=postgres dbname=postgres password=123 sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}
