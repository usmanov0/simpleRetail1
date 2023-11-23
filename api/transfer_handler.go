package api

import (
	"golang-project-template/internal/model"
	"golang-project-template/internal/service"
	"net/http"
	"text/template"
)

type TransferHandler struct {
	TransactionService service.TransactionService
	UserService        service.UserService
}

func (c *TransferHandler) BetweenAccountGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("betweenAccounts").Parse("betweenAccountsTemplate")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (c *TransferHandler) RecipientGet(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.New("recipient").Parse("recipientTemplate")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	recipientsList := []model.Recipient{}

	recipient := model.Recipient{}

	err = tmpl.Execute(w, map[string]interface{}{
		"recipientList": recipientsList,
		"recipient":     recipient,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
