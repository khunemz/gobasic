package repository

import "github.com/jmoiron/sqlx"

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{db: db}
}
