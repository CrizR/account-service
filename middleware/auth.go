package middleware

import (
	log "github.com/sirupsen/logrus"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

func AuthRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		opt := option.WithCredentialsFile("../keys/ecclesia-firebase-key.json")
		app, err := firebase.NewApp(ctx.Request().Context(), nil, opt)
		if err != nil {
			log.Printf("Failed to to create firebase app for auth: %v\n", err)
		}
		auth, err := app.Auth(ctx.Request().Context())
		if err != nil {
			log.Printf("Failed to initialize Auth: %v\n", err)
		}
		_, err = auth.VerifyIDToken(ctx.Request().Context(), ctx.Request().FormValue("firebase_auth_token"))
		if err != nil {
			log.Printf("Failed to verify token: %v\n", err)
		}

		return next(ctx)
	}
}
