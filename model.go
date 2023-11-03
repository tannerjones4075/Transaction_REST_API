package main

import (
	"database/sql"
	"errors"
)

type transaction struct {
	ID       int     `json:"id"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Date     string  `json:"date"`
}

func (p *transaction) getTransaction(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *transaction) updateTransaction(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *transaction) deleteTransaction(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *transaction) createTransaction(db *sql.DB) error {
	return errors.New("Not implemented")
}
func getTransactions(db *sql.DB, start, count int) ([]transaction, error) {
	return nil, errors.New("Not implemented")
}
