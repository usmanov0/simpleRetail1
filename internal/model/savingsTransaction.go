package model

import "time"

type SavingsTransaction struct {
	Id               int       `json:"id"`
	Date             time.Time `json:"date"`
	Description      string    `json:"description"`
	Types            string    `json:"types"`
	Status           string    `json:"status"`
	Amount           float64   `json:"amount"`
	AvailableBalance float64   `json:"availableBalance"`
	SavingsAccountID uint
}
