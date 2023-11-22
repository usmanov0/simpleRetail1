package service

import (
	"context"
	"database/sql"
	"golang-project-template/internal/model"
)

type Services struct {
	Users        UserService
	Accounts     AccountService
	Transactions TransactionService
	Recipient    RecipientService
}

func NewServices(db *sql.DB) *Services {
	return &Services{
		Users:        model.NewUserModel(db),
		Accounts:     model.NewAccountModel(db),
		Transactions: model.NewTransactionModel(db),
		Recipient:    model.NewRecipient(db),
	}
}

type UserService interface {
	CreateUser(ctx context.Context, user *model.Users) error
	UpdateUser(ctx context.Context, user *model.Users) error
	DeleteUser(ctx context.Context, id int) error
}

type AccountService interface {
	CreateAccount(ctx context.Context, account *model.Account) error
	UpdateAccount(ctx context.Context, account *model.Account) error
	DeleteAccount(ctx context.Context, id int) error
}

type TransactionService interface {
	CreateTransaction(ctx context.Context, transaction *model.Transaction) error
	GetTransaction(ctx context.Context, transactionId int) (*model.Transaction, error)
}

type RecipientService interface {
	CreateRecipient(ctx context.Context, recipient *model.Recipient) error
	GetRecipient(ctx context.Context, recipientId int) (*model.Recipient, error)
}
