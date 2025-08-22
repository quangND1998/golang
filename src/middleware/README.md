# Package Middleware

Package middleware cung cấp các middleware cho ứng dụng Fiber, bao gồm middleware logging tích hợp với package logger.

## Tính năng

- ✅ Logging middleware với tích hợp logger
- ✅ Hỗ trợ logging theo level dựa trên HTTP status code
- ✅ Hỗ trợ custom logger instance
- ✅ Hỗ trợ context binding
- ✅ Log đầy đủ thông tin request (method, path, status, duration, IP, User-Agent)

## Sử dụng

### Import

```go
import "github.com/nextlend-api-web-frontend/src/middleware"
```

### 1. Logging Middleware cơ bản

```go
app := fiber.New()

// Thêm middleware logging
app.Use(middleware.LoggingMiddleware())

app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})
```

### 2. Logging Middleware với custom logger

```go
import "github.com/nextlend-api-web-frontend/src/common/logger"

// Tạo custom logger
customLogger := logger.New()

// Sử dụng middleware với custom logger
app.Use(middleware.LoggingMiddlewareWithLogger(customLogger))
```

### 3. Logging Middleware với context

```go
// Middleware sẽ tự động bind context của request
app.Use(middleware.LoggingMiddlewareWithContext())
```

## Log Levels

Middleware sẽ tự động chọn level logging dựa trên HTTP status code:

- **Status >= 500**: `Error` level - Lỗi server
- **Status >= 400**: `Warn` level - Lỗi client (4xx)
- **Status < 400**: `Info` level - Request thành công

## Thông tin được log

Mỗi request sẽ được log với các thông tin sau:

- `method`: HTTP method (GET, POST, PUT, DELETE, etc.)
- `path`: Request path
- `status`: HTTP status code
- `duration`: Thời gian xử lý request
- `ip`: IP address của client
- `user_agent`: User-Agent header
- `error`: Error object (nếu có, chỉ cho status >= 500)

## Ví dụ output

```
2024-01-15T10:30:45.123Z INFO HTTP Request method=GET path=/api/users status=200 duration=15.2ms ip=192.168.1.100 user_agent=Mozilla/5.0...
2024-01-15T10:30:46.456Z WARN HTTP Request Warning method=POST path=/api/login status=401 duration=8.5ms ip=192.168.1.101 user_agent=PostmanRuntime/7.32.3
2024-01-15T10:30:47.789Z ERROR HTTP Request Error method=GET path=/api/data status=500 duration=45.1ms ip=192.168.1.102 user_agent=curl/7.68.0 error=database connection failed
```

## API Reference

### Functions

- `LoggingMiddleware()` - Middleware logging cơ bản sử dụng global logger
- `LoggingMiddlewareWithLogger(logger Logger)` - Middleware với logger instance tùy chỉnh
- `LoggingMiddlewareWithContext()` - Middleware với context binding

## Dependencies

Package này phụ thuộc vào:
- `github.com/gofiber/fiber/v2`
- `github.com/nextlend-api-web-frontend/src/common/logger`

## License

MIT License

