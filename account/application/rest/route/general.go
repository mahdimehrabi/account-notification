package route

import (
	gc "github.com/mahdimehrabi/account-notification/account/application/rest/controller/general"
	"net/http"
)

type General struct {
	mux               *http.ServeMux
	generalController *gc.General
}

func NewGeneral(mux *http.ServeMux) General {
	return General{
		mux:               mux,
		generalController: gc.NewGeneral(),
	}
}

func (g *General) Handle() {
	g.mux.HandleFunc("/api/ping", g.generalController.Ping)
}
