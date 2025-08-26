package middleware

import (
	"time"

	"nextlend-api-web-frontend/src/common/logger"

	"github.com/gofiber/fiber/v2"
)

// LoggingMiddleware tạo middleware logging cho Fiber - phiên bản tối ưu cho performance
func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Thời gian bắt đầu request
		start := time.Now()

		// Thực hiện request
		err := c.Next()

		// Thời gian kết thúc request
		duration := time.Since(start)

		// Lấy thông tin request cơ bản
		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()
		ip := c.IP()

		// Chỉ log errors và warnings, bỏ qua successful requests để tăng performance
		switch {
		case status >= 500:
			logger.Errorw("HTTP Request Error",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				"error", err,
			)
		case status >= 400:
			logger.Warnw("HTTP Request Warning",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
			)
		// Bỏ qua logging cho successful requests (200-399) để tăng performance
		}

		return err
	}
}

// LoggingMiddlewareVerbose tạo middleware logging chi tiết (chỉ dùng cho development)
func LoggingMiddlewareVerbose() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Thời gian bắt đầu request
		start := time.Now()

		// Thực hiện request
		err := c.Next()

		// Thời gian kết thúc request
		duration := time.Since(start)

		// Lấy thông tin request
		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()
		ip := c.IP()

		// Log theo level tùy theo status code
		switch {
		case status >= 500:
			logger.Errorw("HTTP Request Error",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				"error", err,
			)
		case status >= 400:
			logger.Warnw("HTTP Request Warning",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
			)
		default:
			logger.Infow("HTTP Request",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
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
		duration := time.Since(start)

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
		duration := time.Since(start)

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

// LoggingMiddlewareProduction tạo middleware logging tối ưu cho production
// Chỉ log errors và warnings, với sampling cho successful requests
func LoggingMiddlewareProduction(sampleRate float64) fiber.Handler {
	if sampleRate <= 0 || sampleRate > 1 {
		sampleRate = 0.01 // 1% sampling mặc định
	}
	
	return func(c *fiber.Ctx) error {
		// Thời gian bắt đầu request
		start := time.Now()

		// Thực hiện request
		err := c.Next()

		// Thời gian kết thúc request
		duration := time.Since(start)

		// Lấy thông tin request cơ bản
		method := c.Method()
		path := c.Path()
		status := c.Response().StatusCode()
		ip := c.IP()

		// Luôn log errors và warnings
		switch {
		case status >= 500:
			logger.Errorw("HTTP Request Error",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
				"error", err,
			)
		case status >= 400:
			logger.Warnw("HTTP Request Warning",
				"method", method,
				"path", path,
				"status", status,
				"duration", duration,
				"ip", ip,
			)
		default:
			// Sampling cho successful requests để giảm overhead
			if duration > 100*time.Millisecond || // Log slow requests
			   (method != "GET" && method != "HEAD") || // Log non-GET requests
			   (sampleRate > 0 && float64(time.Now().UnixNano()%10000)/10000 < sampleRate) { // Random sampling
				logger.Infow("HTTP Request",
					"method", method,
					"path", path,
					"status", status,
					"duration", duration,
					"ip", ip,
				)
			}
		}

		return err
	}
}

// LoggingMiddlewareMinimal tạo middleware logging tối thiểu - chỉ log errors
func LoggingMiddlewareMinimal() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Thực hiện request
		err := c.Next()

		// Chỉ log errors
		status := c.Response().StatusCode()
		if status >= 500 {
			logger.Errorw("HTTP Request Error",
				"method", c.Method(),
				"path", c.Path(),
				"status", status,
				"ip", c.IP(),
				"error", err,
			)
		}

		return err
	}
}

