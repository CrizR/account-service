package server

import (
	"os"

	"github.com/ecclesia-dev/account-service/controllers"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

type Server struct {
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
	s.echo.POST("/api/create/{userId}", s.createUser)
	s.echo.GET("/api/search/users/{userId}", s.getUserByID)
	s.echo.GET("/api/search/users", s.getAllUsers)
	s.echo.GET("/api/search/email/{email}", s.getUserByEmail)
	s.echo.DELETE("/api/remove/user/{userId}", s.removeUserByID)
	s.echo.PUT("/api/update/user/{userId}", s.updateUser)
}
