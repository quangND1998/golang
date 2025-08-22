package logger

import (
	"context"
	"fmt"
)

// TestExample minh họa cách sử dụng logger
func TestExample() {
	// Khởi tạo logger
	Init()

	// Test các level cơ bản
	Info("Ứng dụng đã khởi động thành công")
	Debug("Debug message cho developer")
	Warn("Cảnh báo: Database connection chậm")
	Error("Lỗi kết nối database")

	// Test format
	Infof("Server đang chạy trên port %d", 8080)
	Errorf("Lỗi xử lý request: %s", "timeout")

	// Test key-value pairs
	Infow("User đăng nhập thành công",
		"user_id", "12345",
		"email", "user@example.com",
		"ip", "192.168.1.1",
	)

	Errorw("Lỗi xử lý payment",
		"transaction_id", "txn_123",
		"amount", 1000000,
		"error", "insufficient_funds",
	)

	// Test với context
	ctx := context.WithValue(context.Background(), "request_id", "req_123")
	loggerWithCtx := WithContext(ctx)
	loggerWithCtx.Info("Request được xử lý")

	// Test custom logger
	customLogger := New()
	customLogger.Info("Custom logger message")

	// Test custom config
	config := &LoggerConfig{
		Level:  LevelDebug,
		Prefix: "[CUSTOM]",
	}
	customLoggerWithConfig := NewWithConfig(config)
	customLoggerWithConfig.Debug("Debug message từ custom logger")

	fmt.Println("Logger test completed successfully!")
}

