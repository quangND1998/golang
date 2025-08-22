package middleware

import (
	"time"

	"nextlend-api-web-frontend/src/common/logger"

	"github.com/gofiber/fiber/v2"
)

// LoggingMiddleware tạo middleware logging cho Fiber
func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Thời gian bắt đầu request
		start := time.Now()

		// Thực hiện request
		err := c.Next()

		// Thời gian kết thúc request
		end := time.Now()
		duration := end.Sub(start)

		// Lấy thông tin request
		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()
		logger.Error("status: ", status)
		ip := c.IP()
		// userAgent := c.Get("User-Agent")

		// Log theo level tùy theo status code
		switch {
		case status >= 500:
			logger.Errorw("HTTP Request Error",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				// "user_agent", userAgent,
				"error", err,
				"request", c.Request().Body(),
				"response", c.Response().Body(),
			)
		case status >= 400:
			logger.Warnw("HTTP Request Warning",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				// "user_agent", userAgent,
				"request", c.Request().Body(),
				"response", c.Response().Body(),
			)
		default:
			logger.Infow("HTTP Request",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				// "user_agent", userAgent,
				"query", c.Queries(),
				"body", string(c.Body()),
				"response", string(c.Response().Body()),
			)
		}

		return err
	}
}

// LoggingMiddlewareWithLogger tạo middleware logging với logger tùy chỉnh
func LoggingMiddlewareWithLogger(log logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Thời gian bắt đầu request
		start := time.Now()

		// Thực hiện request
		err := c.Next()
		// Thời gian kết thúc request
		end := time.Now()
		duration := end.Sub(start)

		// Lấy thông tin request
		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()

		if err != nil {
			if e, ok := err.(*fiber.Error); ok {
				status = e.Code
			} else {
				status = 500
			}
		}
		ip := c.IP()
		userAgent := c.Get("User-Agent")

		// Log theo level tùy theo status code
		switch {
		case status >= 500:
			log.Errorw("HTTP Request Error",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				"user_agent", userAgent,
				"error", err,
			)
		case status >= 400:
			log.Warnw("HTTP Request Warning",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				"user_agent", userAgent,
			)
		default:
			log.Infow("HTTP Request",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				"user_agent", userAgent,
			)
		}

		return err
	}
}

// LoggingMiddlewareWithContext tạo middleware logging với context
func LoggingMiddlewareWithContext() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Thời gian bắt đầu request
		start := time.Now()

		// Thực hiện request
		err := c.Next()

		// Thời gian kết thúc request
		end := time.Now()
		duration := end.Sub(start)

		// Lấy thông tin request
		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()
		ip := c.IP()
		userAgent := c.Get("User-Agent")

		// Tạo logger với context
		log := logger.WithContext(c.Context())

		// Log theo level tùy theo status code
		switch {
		case status >= 500:
			log.Errorw("HTTP Request Error",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				"user_agent", userAgent,
				"error", err,
			)
		case status >= 400:
			log.Warnw("HTTP Request Warning",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				"user_agent", userAgent,
			)
		default:
			log.Infow("HTTP Request",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				"user_agent", userAgent,
			)
		}

		return err
	}
}

