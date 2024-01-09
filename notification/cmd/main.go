package main

import (
	"github.com/mahdimehrabi/account-notification/notification/application/socket"
	infrastructures "github.com/mahdimehrabi/account-notification/notification/internal/infrastructure"
)

func main() {
	env := infrastructures.NewEnv()
	env.LoadEnv()
	socket.RunServer(env)
}
