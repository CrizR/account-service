package main

import (
	"github.com/ecclesia-dev/account-service/controllers"
	"github.com/ecclesia-dev/account-service/server"
)

var bindTo string

func main() {
	acctCtlr := controllers.NewAccountController()
	serv := server.New()
	serv.Start()
}
