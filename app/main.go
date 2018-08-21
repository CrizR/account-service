package main

import (
	"github.com/ecclesia-dev/account-service/data"
	"github.com/ecclesia-dev/account-service/server"
)

func main() {
	accounts := data.NewAccountAccess()
	serv := server.New(accounts)
	serv.Start(":8080")
}
