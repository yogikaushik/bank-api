package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yogikaushik/bank-api/service"

	"github.com/gorilla/mux"
)

// AccountController handles account-related requests
type AccountController struct {
	Service service.AccountService
}

// NewAccountController creates a new AccountController
func NewAccountController(s service.AccountService) *AccountController {
	return &AccountController{Service: s}
}

type AccountInput struct {
	DocumentNumber string `json:"document_number"`
}

// CreateAccount godoc
// @Summary Create a new account
// @Description Create a new account with the given document number
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body AccountInput true "Account document number"
// @Success 200 {object} models.Account
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /accounts [post]
func (c *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var input AccountInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := c.Service.CreateAccount(input.DocumentNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}

// GetAccount godoc
// @Summary Get account information
// @Description Get information of an account by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path uint true "Account ID"
// @Success 200 {object} models.Account
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Account not found"
// @Router /accounts/{id} [get]
func (c *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := c.Service.GetAccount(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(account)
}
