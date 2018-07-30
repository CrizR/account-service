package server

import (
	log "github.com/sirupsen/logrus"
	"github.com/labstack/echo"
	"time"
	"github.com/google/uuid"
)


func (s *Server) setupRequest(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()
		// Default fields
		logger := s.log.WithFields(log.Fields{
			"method":     req.Method,
			"path":       req.URL.Path,
			"request_id": uuid.NewRandom(),
		})
		ctx.Set(loggerKey, logger)
		startTime := time.Now()

		defer func() {
			rsp := ctx.Response()
			// End of Request
			logger.WithFields(log.Fields{
				"status_code":  rsp.Status,
				"runtime_nano": time.Since(startTime).Nanoseconds(),
			}).Info("Finished request")
		}()

		// Starting Request
		logger.WithFields(log.Fields{
			"user_agent":     req.UserAgent(),
			"content_length": req.ContentLength,
		}).Info("Starting request")
		err := f(ctx)
		if err != nil {
			ctx.Error(err)
		}
		return err
	}
}
