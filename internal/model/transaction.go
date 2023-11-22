package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Transaction struct {
	Id               int       `json:"id"`
	Date             time.Time `json:"date"`
	Description      string    `json:"description"`
	Types            string    `json:"types"`
	Status           string    `json:"status"`
	Amount           float64   `json:"amount"`
	AvailableBalance float64   `json:"availableBalance"`
}

func (t *TransactionModel) CreateTransaction(ctx context.Context, transaction *Transaction) error {
	query := `
		INSERT INTO transactions (date, description, types, status, amount, available_balance)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, date, description, types, status, amount, available_balance;
	`
	err := t.db.QueryRowContext(ctx, query,
		transaction.Date,
		transaction.Description,
		transaction.Types,
		transaction.Status,
		transaction.Amount,
		transaction.AvailableBalance,
	).Scan(
		&transaction.Id,
		&transaction.Date,
		&transaction.Description,
		&transaction.Status,
		&transaction.Amount,
		&transaction.AvailableBalance,
		&transaction.Types,
	)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionModel) GetTransaction(ctx context.Context, transactionID int) (*Transaction, error) {
	query := `SELECT * from transactions
	WHERE id=$1`

	transaction := &Transaction{}
	err := t.db.QueryRowContext(ctx, query, transactionID).Scan(
		&transaction.Id,
		&transaction.Date,
		&transaction.Description,
		&transaction.Status,
		&transaction.Amount,
		&transaction.AvailableBalance,
		&transaction.Types)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, fmt.Errorf("Transaction not found: %w,", err)
		}
	}
	return transaction, nil
}
