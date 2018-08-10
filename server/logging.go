package server

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func (s *Server) LogRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()
		// Default fields

		rand, _ := uuid.NewRandom()

		logger := s.log.WithFields(log.Fields{
			"method":     req.Method,
			"path":       req.URL.Path,
			"request_id": rand,
		})
		ctx.Set(rand.String(), logger)
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

		return next(ctx)
	}
}
