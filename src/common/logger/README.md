# Package Logger

Package logger cung cấp một wrapper cho Fiber log với các tính năng mở rộng và dễ sử dụng.

## Tính năng

- ✅ Tích hợp với Fiber log
- ✅ Hỗ trợ nhiều level logging (Trace, Debug, Info, Warn, Error, Fatal, Panic)
- ✅ Hỗ trợ format string và key-value pairs
- ✅ Hỗ trợ context binding
- ✅ Middleware cho Fiber
- ✅ Ghi log vào file, console hoặc cả hai
- ✅ Log theo ngày
- ✅ Cấu hình linh hoạt

## Cài đặt

Package này sử dụng Fiber log, đảm bảo bạn đã có dependency:

```go
go get github.com/gofiber/fiber/v2
```

## Sử dụng cơ bản

### 1. Khởi tạo logger

```go
import "your-project/src/common/logger"

// Khởi tạo với cấu hình mặc định
logger.Init()

// Hoặc tạo logger instance riêng
log := logger.New()
```

### 2. Logging cơ bản

```go
// Log các level khác nhau
logger.Info("Ứng dụng đã khởi động")
logger.Debug("Debug message")
logger.Warn("Cảnh báo")
logger.Error("Lỗi xảy ra")
```

// Log với format
logger.Infof("Server đang chạy trên port %d", 8080)
logger.Errorf("Lỗi: %s", err.Error())

// Log với key-value pairs
logger.Infow("User đăng nhập",
    "user_id", "12345",
    "email", "user@example.com",
    "ip", "192.168.1.1",
)
```

### 3. Sử dụng với context

```go
ctx := context.WithValue(context.Background(), "request_id", "req_123")
loggerWithCtx := logger.WithContext(ctx)
loggerWithCtx.Info("Request được xử lý")
```

## Cấu hình nâng cao

### 1. Tạo logger với cấu hình tùy chỉnh

```go
config := &logger.LoggerConfig{
    Level:      logger.LevelDebug,
    TimeFormat: "2006-01-02 15:04:05",
    Prefix:     "[NEXTLEND-API]",
}

customLogger := logger.NewWithConfig(config)
customLogger.Debug("Debug message")
```

### 2. Ghi log vào file

```go
// Ghi vào file đơn
err := logger.InitFileLogger("app.log")
if err != nil {
    log.Fatal(err)
}

// Ghi vào cả console và file
err = logger.InitMultiOutputLogger("app.log")
if err != nil {
    log.Fatal(err)
}

// Ghi theo ngày
err = logger.InitDailyLogger("./logs")
if err != nil {
    log.Fatal(err)
}
```

### 3. Thiết lập level

```go
// Chỉ log từ Info trở lên
logger.SetLevel(logger.LevelInfo)

// Log tất cả (bao gồm Debug, Trace)
logger.SetLevel(logger.LevelTrace)
```

## Middleware cho Fiber

### 1. Middleware cơ bản

```go
import "github.com/nextlend-api-web-frontend/src/middleware"

app := fiber.New()

// Thêm middleware logging
app.Use(middleware.LoggingMiddleware())

app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})
```

### 2. Middleware với logger tùy chỉnh

```go
import "github.com/nextlend-api-web-frontend/src/middleware"

customLogger := logger.New()
app.Use(middleware.LoggingMiddlewareWithLogger(customLogger))
```

### 3. Middleware với context

```go
app.Use(middleware.LoggingMiddlewareWithContext())
```

## Ví dụ hoàn chỉnh

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "your-project/src/common/logger"
    "your-project/src/middleware"
)

func main() {
    // Khởi tạo logger
    logger.Init()
    
    // Tạo Fiber app
    app := fiber.New()
    
    // Thêm middleware logging
    app.Use(middleware.LoggingMiddleware())
    
    // Routes
    app.Get("/", func(c *fiber.Ctx) error {
        logger.Info("Homepage được truy cập")
        return c.SendString("Hello, World!")
    })
    
    app.Get("/users/:id", func(c *fiber.Ctx) error {
        userID := c.Params("id")
        logger.Infow("User profile được truy cập",
            "user_id", userID,
            "ip", c.IP(),
        )
        return c.JSON(fiber.Map{"user_id": userID})
    })
    
    // Khởi động server
    logger.Info("Server đang khởi động trên port 3000")
    app.Listen(":3000")
}
```

## Log Levels

| Level | Mô tả |
|-------|-------|
| `LevelTrace` | Chi tiết nhất, thường chỉ dùng khi debug |
| `LevelDebug` | Thông tin debug cho developer |
| `LevelInfo` | Thông tin chung về hoạt động của ứng dụng |
| `LevelWarn` | Cảnh báo, có thể gây vấn đề |
| `LevelError` | Lỗi xảy ra nhưng không làm crash ứng dụng |
| `LevelFatal` | Lỗi nghiêm trọng, ứng dụng sẽ thoát |
| `LevelPanic` | Lỗi nghiêm trọng, gây panic |

## Best Practices

1. **Sử dụng level phù hợp**: Chỉ log những gì cần thiết
2. **Structured logging**: Sử dụng key-value pairs thay vì string concatenation
3. **Context**: Sử dụng context để track request flow
4. **File rotation**: Sử dụng daily logger để tránh file log quá lớn
5. **Performance**: Tránh log quá nhiều trong production

## API Reference

### Functions

- `Init()` - Khởi tạo global logger
- `InitWithConfig(config *LoggerConfig)` - Khởi tạo với cấu hình tùy chỉnh
- `InitFileLogger(filename string)` - Khởi tạo logger ghi vào file
- `InitMultiOutputLogger(filename string)` - Khởi tạo logger ghi vào cả console và file
- `InitDailyLogger(logDir string)` - Khởi tạo logger ghi theo ngày

### Global Functions

- `Info(msg string, args ...interface{})`
- `Infof(format string, args ...interface{})`
- `Infow(msg string, keysAndValues ...interface{})`
- `Debug(msg string, args ...interface{})`
- `Debugf(format string, args ...interface{})`
- `Debugw(msg string, keysAndValues ...interface{})`
- `Warn(msg string, args ...interface{})`
- `Warnf(format string, args ...interface{})`
- `Warnw(msg string, keysAndValues ...interface{})`
- `Error(msg string, args ...interface{})`
- `Errorf(format string, args ...interface{})`
- `Errorw(msg string, keysAndValues ...interface{})`
- `Fatal(msg string, args ...interface{})`
- `Fatalf(format string, args ...interface{})`
- `Fatalw(msg string, keysAndValues ...interface{})`
- `Panic(msg string, args ...interface{})`
- `Panicf(format string, args ...interface{})`
- `Panicw(msg string, keysAndValues ...interface{})`

### Middleware (trong package middleware)

- `middleware.LoggingMiddleware()` - Middleware logging cơ bản
- `middleware.LoggingMiddlewareWithLogger(logger Logger)` - Middleware với logger tùy chỉnh
- `middleware.LoggingMiddlewareWithContext()` - Middleware với context

## License

MIT License
