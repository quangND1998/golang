package database

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var pool = make(map[string]*gorm.DB)

// ConnectAllFromEnv load tất cả connections từ ENV và kết nối
func ConnectAllFromEnv() {
	connections := GetConnections()

	log.Info("🔍 Tìm thấy %d database connections", len(connections))

	for connectionName, config := range connections {
		log.Info("🔗 Kết nối database [%s] (%s)...", connectionName, config.Driver)

		// Validate cấu hình trước khi kết nối
		if err := validateDBConfig(connectionName, config); err != nil {
			log.Info("⚠️ Bỏ qua connection %s: %v", connectionName, err)
			continue
		}

		// Xử lý đặc biệt cho MongoDB
		if config.Driver == "mongodb" {
			log.Info("⚠️ MongoDB [%s] chưa được implement trong GORM, bỏ qua", connectionName)
			continue
		}

		db, err := connectDB(connectionName, config)
		if err != nil {
			log.Info("❌ Lỗi kết nối database %s: %v", connectionName, err)
			continue
		}

		pool[connectionName] = db
		log.Info("✅ Kết nối database [%s] (%s) thành công", connectionName, config.Driver)
	}

	if len(pool) == 0 {
		log.Info("⚠️ Không có database nào được kết nối thành công")
	} else {
		log.Info("📊 Tổng cộng %d database đã được kết nối", len(pool))
	}
}

// connectDB mở kết nối với connection config
func connectDB(connectionName string, config ConnectionConfig) (*gorm.DB, error) {
	dsn := buildDSN(config)
	if dsn == "" {
		return nil, fmt.Errorf("không thể xây dựng DSN cho connection: %s", connectionName)
	}
	fmt.Println("Connecting to database with DSN:", dsn)
	var dialector gorm.Dialector
	switch config.Driver {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", config.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("không thể mở kết nối database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("không thể lấy underlying DB: %v", err)
	}

	// Cấu hình connection pool
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

	// Ping test
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping database error: %v", err)
	}

	return db, nil
}

// Get trả về *gorm.DB theo tên connection
func Get(connectionName string) *gorm.DB {
	db, ok := pool[connectionName]
	if !ok {
		log.Fatalf("❌ Database connection [%s] chưa được cấu hình", connectionName)
	}
	return db
}

// GetSafe trả về *gorm.DB theo tên connection, không panic nếu không tìm thấy
func GetSafe(connectionName string) (*gorm.DB, bool) {
	db, ok := pool[connectionName]
	return db, ok
}

// ListConnections trả về danh sách tên các connection đã kết nối
func ListConnections() []string {
	names := make([]string, 0, len(pool))
	for name := range pool {
		names = append(names, name)
	}
	return names
}

// GetConnectionConfig trả về cấu hình của một connection
func GetConnectionConfig(connectionName string) (ConnectionConfig, bool) {
	connections := GetConnections()
	config, ok := connections[connectionName]
	return config, ok
}

// GetAllConnectionConfigs trả về tất cả cấu hình connections
func GetAllConnectionConfigs() map[string]ConnectionConfig {
	return GetConnections()
}

// IsConnected kiểm tra xem một connection đã được kết nối chưa
func IsConnected(connectionName string) bool {
	_, ok := pool[connectionName]
	return ok
}

// CloseConnection đóng một connection cụ thể
func CloseConnection(connectionName string) error {
	db, ok := pool[connectionName]
	if !ok {
		return fmt.Errorf("connection [%s] không tồn tại", connectionName)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("không thể lấy underlying DB: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("lỗi đóng connection [%s]: %v", connectionName, err)
	}

	delete(pool, connectionName)
	log.Info("🔌 Đã đóng connection [%s]", connectionName)
	return nil
}

// CloseAllConnections đóng tất cả connections
func CloseAllConnections() {
	for connectionName := range pool {
		CloseConnection(connectionName)
	}
	log.Info("🔌 Đã đóng tất cả database connections")
}
