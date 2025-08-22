package database

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Init khởi tạo tất cả database connections
func Init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Không tìm thấy file .env, dùng ENV hệ thống")
	} else {
		log.Println("✅ Đã load file .env thành công")
	}

	// Log thông tin debug về các biến môi trường database
	logDBEnvVars()

	// Kết nối tất cả database connections
	ConnectAllFromEnv()

	// Kết nối MongoDB databases (nếu có)
	// ConnectMongoFromEnv()
}

// logDBEnvVars log các biến môi trường database để debug
func logDBEnvVars() {
	log.Println("🔍 Kiểm tra các biến môi trường database:")

	// Tìm tất cả biến môi trường liên quan đến database
	dbVars := make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		key := parts[0]
		value := parts[1]

		// Kiểm tra các biến môi trường database
		if isDatabaseEnvVar(key) {
			// Ẩn giá trị nhạy cảm
			if strings.Contains(key, "PASSWORD") || strings.Contains(key, "URI") || strings.Contains(key, "URL") {
				if len(value) > 10 {
					value = value[:10] + "..."
				}
			}
			dbVars[key] = value
		}
	}

	if len(dbVars) == 0 {
		log.Println("   Không tìm thấy biến môi trường database nào")
	} else {
		log.Println("   Các biến môi trường database được tìm thấy:")
		for key, value := range dbVars {
			log.Printf("   %s = %s", key, value)
		}
	}
}

// isDatabaseEnvVar kiểm tra xem biến môi trường có liên quan đến database không
func isDatabaseEnvVar(key string) bool {
	databasePrefixes := []string{
		"DB_", "DATABASE_", "APP_DB_", "DATA_DB_", "LOGS_DB_",
		"MONGODB_", "MYSQL_", "POSTGRES_", "SQLITE_", "SQLSRV_",
	}

	for _, prefix := range databasePrefixes {
		if strings.HasPrefix(key, prefix) {
			return true
		}
	}

	return false
}

// ConnectMongoFromEnv kết nối MongoDB databases từ ENV
// Hàm này được implement trong mongo.go

// maskSensitiveInfo ẩn thông tin nhạy cảm trong log
func maskSensitiveInfo(value string) string {
	if value == "" {
		return ""
	}

	if len(value) > 20 {
		return value[:10] + "..." + value[len(value)-10:]
	}

	return value
}

// GetDefaultConnection trả về connection mặc định
func GetDefaultConnection() string {
	// Ưu tiên theo thứ tự: mysql, pgsql, sqlite
	connections := ListConnections()

	priority := []string{"mysql", "pgsql", "sqlite"}
	for _, conn := range priority {
		for _, connected := range connections {
			if connected == conn {
				return conn
			}
		}
	}

	// Nếu không có connection ưu tiên, trả về connection đầu tiên
	if len(connections) > 0 {
		return connections[0]
	}

	return ""
}

// GetConnectionInfo trả về thông tin chi tiết của một connection
func GetConnectionInfo(connectionName string) map[string]interface{} {
	config, exists := GetConnectionConfig(connectionName)
	if !exists {
		return nil
	}

	info := map[string]interface{}{
		"driver":    config.Driver,
		"host":      config.Host,
		"port":      config.Port,
		"database":  config.Database,
		"username":  config.Username,
		"charset":   config.Charset,
		"prefix":    config.Prefix,
		"strict":    config.Strict,
		"connected": IsConnected(connectionName),
	}

	// Thêm thông tin pool nếu đã kết nối
	if IsConnected(connectionName) {
		info["max_open_conns"] = config.MaxOpenConns
		info["max_idle_conns"] = config.MaxIdleConns
		info["conn_max_lifetime"] = config.ConnMaxLifetime.String()
	}

	return info
}

// GetAllConnectionsInfo trả về thông tin tất cả connections
func GetAllConnectionsInfo() map[string]map[string]interface{} {
	connections := GetConnections()
	info := make(map[string]map[string]interface{})

	for connectionName := range connections {
		info[connectionName] = GetConnectionInfo(connectionName)
	}

	return info
}
