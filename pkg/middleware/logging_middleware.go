package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shardeum/service-validator/pkg/constants"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()
			deviceId := c.Request().Host
			HttpMiddlewareAccessLogger(req.Method, req.RequestURI, res.Status, res.Size, time.Since(start), deviceId)
			return nil
		}
	}
}

func HttpMiddlewareAccessLogger(method string, uri string, status int, size int64, time time.Duration, deviceId string) {
	log.WithFields(log.Fields{
		constants.METHOD:    method,
		constants.URI:       uri,
		constants.STATUS:    status,
		constants.SIZE:      size,
		constants.TIME:      time,
		constants.DEVICE_ID: deviceId,
	}).Info(
		constants.HTTP,
	)
}
