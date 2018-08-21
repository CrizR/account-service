package server

import (
	"github.com/ecclesia-dev/account-service/controllers"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Server struct {
	log      *logrus.Entry
	echo     *echo.Echo
	accounts controllers.AccountController
}

func New(ctlr controllers.AccountController) Server {
	e := echo.New()
	e.HideBanner = true
	logger := logrus.NewEntry(logrus.New())
	serv := Server{log: logger, echo: e, accounts: ctlr}
	e.Use(serv.LogRequest)
	return serv
}

func (s *Server) Start(port string) {
	s.setRoutes()
	s.echo.Logger.Fatal(s.echo.Start(port))
}

func (s *Server) setRoutes() {
	s.echo.POST("/api/accounts/create", s.createAccount)
	s.echo.GET("/api/accounts/id=:id", s.getAccountByID)
	s.echo.GET("/api/accounts/login/username=:email&password=:password", s.login)
	s.echo.GET("/api/accounts/logout/id=:id", s.logout)
	s.echo.GET("/api/accounts/oauth/refresh/token/id=:id", s.RefreshToken)
	s.echo.GET("/api/accounts", s.getAllAccounts)
	s.echo.GET("/api/accounts/email/email=:email", s.getAccountByEmail)
	s.echo.PUT("/api/accounts/id=:id", s.updateAccount)
}
