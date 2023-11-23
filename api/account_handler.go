package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang-project-template/internal/model"
	"golang-project-template/internal/service"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	AccountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) *AccountHandler {
	return &AccountHandler{AccountService: accountService}
}

func (h *AccountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call service to create account
	if err := h.AccountService.CreateAccount(r.Context(), &account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account := &model.Account{Id: id}
	if err := h.AccountService.UpdateAccount(r.Context(), account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (api *AccountHandler) AllHandlers() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/accounts", api.CreateAccountHandler).Methods("POST")
	router.HandleFunc("/api/accounts/{id:[0-9]+}", api.UpdateAccountHandler).Methods("PUT")
	router.HandleFunc("/api/accounts/{id:[0-9]+}", api.DeleteAccountHandler).Methods("DELETE")
	return router
}

func (api *AccountHandler) DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromRequest(r)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	err = api.AccountService.DeleteAccount(r.Context(), id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusNoContent, nil)
}

func sendError(w http.ResponseWriter, statusCode int, err error) {
	sendResponse(w, statusCode, map[string]string{"error": err.Error()})
}

func sendResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func parseIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return 0, nil
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, nil
	}

	return id, nil
}
