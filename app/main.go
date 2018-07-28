package main

import (
	"github.com/ecclesia-dev/account-service/controllers"
	"github.com/ecclesia-dev/account-service/server"
)

func main() {
	acctCtlr := controllers.NewAccountController()
	serv := server.New(acctCtlr)
	serv.Start(":8080")
}
