package model

import (
	"context"
	"database/sql"
	"fmt"
)

type Recipient struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	AccountNumber int    `json:"accountNumber"`
	Description   string `json:"description"`
}

func (r *RecipientModel) CreateRecipient(ctx context.Context, recipient *Recipient) error {
	query := `INSERT INTO recipients (name, email, phone, account_number, description)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, email, phone, account_number, description;`

	err := r.db.QueryRowContext(ctx, query,
		recipient.Name,
		recipient.Phone,
		recipient.Email,
	).Scan(&recipient.Id,
		&recipient.Name,
		&recipient.Email,
		&recipient.Phone,
		&recipient.AccountNumber,
		&recipient.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *RecipientModel) GetRecipient(ctx context.Context, recipientId int) (*Recipient, error) {
	query := `select name,account_number
	FROM recipients
	WHERE id=$1`

	recipient := &Recipient{}
	err := r.db.QueryRowContext(ctx, query, recipientId).Scan(
		&recipient.Id,
		&recipient.Name,
		&recipient.Email,
		&recipient.Phone,
		&recipient.AccountNumber,
		&recipient.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Recipient not found %w", err)
		}
	}
	return recipient, nil
}
