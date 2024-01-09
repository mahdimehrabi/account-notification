package account

import (
	"encoding/json"
	"errors"
	accountRepo "github.com/mahdimehrabi/account-notification/account/internal/repository/account"
	"github.com/mahdimehrabi/account-notification/account/internal/service"
	"github.com/mahdimehrabi/graph-interview/receiver/application/http/response"
	"net/http"
)

type Account struct {
	accountService *service.AccountService
}

func NewAccount() *Account {
	return &Account{
		accountService: service.NewAccountService(),
	}
}

func (c Account) AccountsHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.List(w, r)
		return
	case http.MethodPost:
		c.Create(w, r)
		return

	}
}

func (c Account) Create(w http.ResponseWriter, r *http.Request) {
	dto := CreateDTO{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dto); err != nil {
		response.BadRequestErrorResponse(w)
		return
	}
	if err := dto.Validate(); err != nil {
		response.BadRequestErrorResponse(w)
		return
	}
	if err := c.accountService.Create(dto.ToModel()); err != nil {
		response.InternalErrorResponse(w)
		return
	}

	response.SuccessResponse(w, nil, "account created successfully")
}

func (c Account) List(w http.ResponseWriter, r *http.Request) {
	accounts, err := c.accountService.List()
	if err != nil {
		response.InternalErrorResponse(w)
	}

	resp, err := json.Marshal(accounts)
	if err != nil {
		response.InternalErrorResponse(w)
		return
	}

	response.SuccessResponse(w, string(resp), "")
}

func (c Account) Send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.MethodNotAllowedErrorResponse(w)
		return
	}

	dto := TransactionDTO{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dto); err != nil {
		response.BadRequestErrorResponse(w)
		return
	}
	if err := dto.Validate(); err != nil {
		response.BadRequestErrorResponse(w)
		return
	}

	if err := c.accountService.Send(dto.ToModel()); err != nil {
		if errors.Is(err, accountRepo.ErrNotFound) || errors.Is(err, accountRepo.ErrInsufficientAmount) {
			response.GenResponse(w, http.StatusBadRequest, nil, err.Error(), nil)
			return
		}
		response.InternalErrorResponse(w)
		return
	}

	response.SuccessResponse(w, nil, `transaction sent successfully`)

}
