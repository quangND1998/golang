package logger

import (
	"context"
	"io"
	"os"
	"time"

	fiberlog "github.com/gofiber/fiber/v2/log"
)

// Logger interface định nghĩa các method logging cần thiết
type Logger interface {
	// Các method cơ bản
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	// Các method format
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	// Các method với key-value pairs
	Tracew(msg string, keysAndValues ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})

	// Method với context
	WithContext(ctx context.Context) Logger
}

// LoggerConfig cấu hình cho logger
type LoggerConfig struct {
	Level      fiberlog.Level
	Output     io.Writer
	TimeFormat string
	Prefix     string
}

// DefaultConfig trả về cấu hình mặc định
func DefaultConfig() *LoggerConfig {
	return &LoggerConfig{
		Level:      fiberlog.LevelInfo,
		Output:     os.Stdout,
		TimeFormat: time.RFC3339,
		Prefix:     "[NEXTLEND]",
	}
}

// fiberLogger implementation của Logger interface sử dụng Fiber log
type fiberLogger struct {
	ctx context.Context
}

// New tạo một logger mới với cấu hình mặc định
func New() Logger {
	return &fiberLogger{}
}

// NewWithConfig tạo một logger mới với cấu hình tùy chỉnh
func NewWithConfig(config *LoggerConfig) Logger {
	if config == nil {
		config = DefaultConfig()
	}

	// Thiết lập level
	fiberlog.SetLevel(config.Level)

	// Thiết lập output
	if config.Output != nil {
		fiberlog.SetOutput(config.Output)
	}

	return &fiberLogger{}
}

// NewFileLogger tạo logger ghi vào file
func NewFileLogger(filename string) (Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	config := DefaultConfig()
	config.Output = file

	return NewWithConfig(config), nil
}

// NewMultiOutputLogger tạo logger ghi vào cả console và file
func NewMultiOutputLogger(filename string) (Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	config := DefaultConfig()
	config.Output = multiWriter

	return NewWithConfig(config), nil
}

// Implementation của các method Trace
func (l *fiberLogger) Trace(args ...interface{}) {
	fiberlog.Trace(args...)
}

func (l *fiberLogger) Tracef(format string, args ...interface{}) {
	fiberlog.Tracef(format, args...)
}

func (l *fiberLogger) Tracew(msg string, keysAndValues ...interface{}) {
	fiberlog.Tracew(msg, keysAndValues...)
}

// Implementation của các method Debug
func (l *fiberLogger) Debug(args ...interface{}) {
	fiberlog.Debug(args...)
}

func (l *fiberLogger) Debugf(format string, args ...interface{}) {
	fiberlog.Debugf(format, args...)
}

func (l *fiberLogger) Debugw(msg string, keysAndValues ...interface{}) {
	fiberlog.Debugw(msg, keysAndValues...)
}

// Implementation của các method Info
func (l *fiberLogger) Info(args ...interface{}) {
	fiberlog.Info(args...)
}

func (l *fiberLogger) Infof(format string, args ...interface{}) {
	fiberlog.Infof(format, args...)
}

func (l *fiberLogger) Infow(msg string, keysAndValues ...interface{}) {
	fiberlog.Infow(msg, keysAndValues...)
}

// Implementation của các method Warn
func (l *fiberLogger) Warn(args ...interface{}) {
	fiberlog.Warn(args...)
}

func (l *fiberLogger) Warnf(format string, args ...interface{}) {
	fiberlog.Warnf(format, args...)
}

func (l *fiberLogger) Warnw(msg string, keysAndValues ...interface{}) {
	fiberlog.Warnw(msg, keysAndValues...)
}

// Implementation của các method Error
func (l *fiberLogger) Error(args ...interface{}) {
	fiberlog.Error(args...)
}

func (l *fiberLogger) Errorf(format string, args ...interface{}) {
	fiberlog.Errorf(format, args...)
}

func (l *fiberLogger) Errorw(msg string, keysAndValues ...interface{}) {
	fiberlog.Errorw(msg, keysAndValues...)
}

// Implementation của các method Fatal
func (l *fiberLogger) Fatal(args ...interface{}) {
	fiberlog.Fatal(args...)
}

func (l *fiberLogger) Fatalf(format string, args ...interface{}) {
	fiberlog.Fatalf(format, args...)
}

func (l *fiberLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	fiberlog.Fatalw(msg, keysAndValues...)
}

// Implementation của các method Panic
func (l *fiberLogger) Panic(args ...interface{}) {
	fiberlog.Panic(args...)
}

func (l *fiberLogger) Panicf(format string, args ...interface{}) {
	fiberlog.Panicf(format, args...)
}

func (l *fiberLogger) Panicw(msg string, keysAndValues ...interface{}) {
	fiberlog.Panicw(msg, keysAndValues...)
}

// WithContext trả về logger với context được bind
func (l *fiberLogger) WithContext(ctx context.Context) Logger {
	return &fiberLogger{ctx: ctx}
}
