package main

import (
	http "github.com/mahdimehrabi/account-notification/account/application/rest"
	infrastructures "github.com/mahdimehrabi/account-notification/account/internal/infrastructure"
)

func main() {
	env := infrastructures.NewEnv()
	env.LoadEnv()

	http.Bootstrap(env)
}
