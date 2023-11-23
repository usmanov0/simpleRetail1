package model

import (
	"context"
	"fmt"
)

type Account struct {
	Id             int     `json:"id"`
	AccountNumber  int     `json:"accountNumber"`
	AccountBalance float64 `json:"accountBalance"`
}

func (a *AccountModel) CreateAccount(ctx context.Context, account *Account) error {
	query := `INSERT INTO accounts (account_number, account_balance)
	VALUES ($1,$2)
	RETURNING id,account_number,account_balance;`

	err := a.db.QueryRowContext(ctx, query, account.AccountNumber, account.AccountBalance).Scan(
		&account.Id, &account.AccountNumber, &account.AccountBalance)
	if err != nil {
		return fmt.Errorf("error creating account: %w", err)
	}
	return nil
}

func (a *AccountModel) UpdateAccount(ctx context.Context, account *Account) error {
	query := `UPDATE account
	SET account_number=$2
	WHERE id=$1
	RETURNING id,account_number`

	return a.db.QueryRowContext(ctx, query, account.Id).Scan(
		&account.Id, &account.AccountNumber, &account.AccountBalance)
}

func (a *AccountModel) DeleteAccount(ctx context.Context, id int) error {
	query := `DELETE from accounts WHERE id=$1`

	_, err := a.db.ExecContext(ctx, query, id)
	return err
}

func (a *AccountModel) Deposit(accountType string, amount float64, principal *Account) {

}

func (a *AccountModel) WithDraw(accountType string, amount float64, principal *Account) {

}
