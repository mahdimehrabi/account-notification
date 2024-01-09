package route

import (
	"github.com/mahdimehrabi/account-notification/account/application/rest/controller/account"
	"net/http"
)

type Account struct {
	mux               *http.ServeMux
	accountController *account.Account
}

func NewAccount(mux *http.ServeMux) Account {
	return Account{
		mux:               mux,
		accountController: account.NewAccount(), //transient dependency injection (improves performance)
	}
}

func (g *Account) Handle() {
	g.mux.HandleFunc("/api/accounts/", g.accountController.AccountsHandle)
	g.mux.HandleFunc("/api/accounts/send/", g.accountController.Send)
}
