package server

import (
	"os"

	"github.com/ecclesia-dev/account-service/controllers"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

type Server struct {
	log    *logrus.Entry
	echo     *echo.Echo
	accounts controllers.AccountController
}

func New(ctlr controllers.AccountController) Server {
	e := echo.New()
	e.HideBanner = true
	logger := mw.LoggerWithConfig(mw.LoggerConfig{
		Skipper: mw.DefaultSkipper,
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency}}` + "\n",
		Output: os.Stdout,
	})
	e.Use(logger)

	serv := Server{echo: e, accounts: ctlr}
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
