package model

import "database/sql"

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db: db}
}

type AccountModel struct {
	db *sql.DB
}

func NewAccountModel(db *sql.DB) *AccountModel {
	return &AccountModel{db: db}
}

type TransactionModel struct {
	db *sql.DB
}

func NewTransactionModel(db *sql.DB) *TransactionModel {
	return &TransactionModel{db: db}
}

type SavingsAccountsModel struct {
	db *sql.DB
}

func NewSavingsAccounts(db *sql.DB) *SavingsAccountsModel {
	return &SavingsAccountsModel{db: db}
}

type SavingsTransactionModel struct {
	db *sql.DB
}

func NewSavingsTransaction(db *sql.DB) *SavingsTransactionModel {
	return &SavingsTransactionModel{db: db}
}

type RecipientModel struct {
	db *sql.DB
}

func NewRecipient(db *sql.DB) *RecipientModel {
	return &RecipientModel{db: db}
}
