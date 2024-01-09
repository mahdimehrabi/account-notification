package http

import (
	"github.com/mahdimehrabi/account-notification/account/application/rest/dpi"
	"github.com/mahdimehrabi/account-notification/account/application/rest/route"
	infrastructures "github.com/mahdimehrabi/account-notification/account/internal/infrastructure"
	"log"
	"net/http"
	"time"
)

func HandleRoutes(mx *http.ServeMux) {
	general := route.NewGeneral(mx)
	general.Handle()

	account := route.NewAccount(mx)
	account.Handle()
}

func Bootstrap(env *infrastructures.Env) {
	dpi.AppDPI = dpi.NewDPI(env)

	mx := http.NewServeMux()
	HandleRoutes(mx)
	server := &http.Server{
		Addr:              ":" + env.ServerPort,
		Handler:           mx,
		ReadHeaderTimeout: time.Second * 5, // prevent Slow-loris attack
	}
	log.Printf("running REST API on port \":%s\" 🏁", env.ServerPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
