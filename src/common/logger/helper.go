package logger

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	fiberlog "github.com/gofiber/fiber/v2/log"
)

var (
	// Global logger instance
	globalLogger Logger
)

// Init khởi tạo global logger với cấu hình mặc định
func Init() {
	globalLogger = New()
}

// InitWithConfig khởi tạo global logger với cấu hình tùy chỉnh
func InitWithConfig(config *LoggerConfig) {
	globalLogger = NewWithConfig(config)
}

// InitFileLogger khởi tạo global logger ghi vào file
func InitFileLogger(filename string) error {
	logger, err := NewFileLogger(filename)
	if err != nil {
		return err
	}
	globalLogger = logger
	return nil
}

// InitMultiOutputLogger khởi tạo global logger ghi vào cả console và file
func InitMultiOutputLogger(filename string) error {
	logger, err := NewMultiOutputLogger(filename)
	if err != nil {
		return err
	}
	globalLogger = logger
	return nil
}

// GetGlobalLogger trả về global logger instance
func GetGlobalLogger() Logger {
	if globalLogger == nil {
		Init()
	}
	return globalLogger
}

// SetGlobalLogger thiết lập global logger
func SetGlobalLogger(logger Logger) {
	globalLogger = logger
}

// Global logging functions - sử dụng global logger
func Trace(args ...interface{}) {
	GetGlobalLogger().Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	GetGlobalLogger().Tracef(format, args...)
}

func Tracew(msg string, keysAndValues ...interface{}) {
	GetGlobalLogger().Tracew(msg, keysAndValues...)
}

func Debug(args ...interface{}) {
	GetGlobalLogger().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	GetGlobalLogger().Debugf(format, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	GetGlobalLogger().Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	GetGlobalLogger().Info(args...)
}

func Infof(format string, args ...interface{}) {
	GetGlobalLogger().Infof(format, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	GetGlobalLogger().Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	GetGlobalLogger().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	GetGlobalLogger().Warnf(format, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	GetGlobalLogger().Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	GetGlobalLogger().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	GetGlobalLogger().Errorf(format, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	GetGlobalLogger().Errorw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	GetGlobalLogger().Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	GetGlobalLogger().Fatalf(format, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	GetGlobalLogger().Fatalw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	GetGlobalLogger().Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	GetGlobalLogger().Panicf(format, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	GetGlobalLogger().Panicw(msg, keysAndValues...)
}

// WithContext trả về global logger với context
func WithContext(ctx context.Context) Logger {
	return GetGlobalLogger().WithContext(ctx)
}

// Helper functions cho việc tạo log file theo ngày
func CreateDailyLogFile(logDir string) (string, error) {
	// Tạo thư mục log nếu chưa tồn tại
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return "", err
	}

	// Tạo tên file theo ngày
	date := time.Now().Format("2006-01-02")
	filename := filepath.Join(logDir, "app-"+date+".log")
	
	return filename, nil
}

// InitDailyLogger khởi tạo logger ghi vào file theo ngày
func InitDailyLogger(logDir string) error {
	filename, err := CreateDailyLogFile(logDir)
	if err != nil {
		return err
	}
	
	return InitMultiOutputLogger(filename)
}

// SetLevel thiết lập level cho global logger
func SetLevel(level fiberlog.Level) {
	fiberlog.SetLevel(level)
}

// SetOutput thiết lập output cho global logger
func SetOutput(w io.Writer) {
	fiberlog.SetOutput(w)
}

// LogLevel constants để dễ sử dụng
const (
	LevelTrace = fiberlog.LevelTrace
	LevelDebug = fiberlog.LevelDebug
	LevelInfo  = fiberlog.LevelInfo
	LevelWarn  = fiberlog.LevelWarn
	LevelError = fiberlog.LevelError
	LevelFatal = fiberlog.LevelFatal
	LevelPanic = fiberlog.LevelPanic
)
