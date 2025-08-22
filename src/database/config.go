package database

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ConnectionConfig cấu hình cho một database connection
type ConnectionConfig struct {
	Driver          string
	URL             string
	Host            string
	Port            string
	Database        string
	Username        string
	Password        string
	Charset         string
	Collation       string
	Prefix          string
	PrefixIndexes   bool
	Strict          bool
	Engine          string
	SSLMode         string
	Schema          string
	Timezone        string
	ParseTime       bool
	Loc             string
	UnixSocket      string
	ForeignKeys     bool
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	Options         map[string]interface{}
}

// getEnv lấy ENV hoặc giá trị mặc định
func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

// getRequiredEnv lấy ENV bắt buộc, panic nếu không có
func getRequiredEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("❌ Biến môi trường %s là bắt buộc", key))
	}
	return val
}

// parseDuration parse chuỗi duration kiểu "30m", "1h"
func parseDuration(str string) time.Duration {
	d, err := time.ParseDuration(str)
	if err != nil {
		return 0
	}
	return d
}

// atoi chuyển string thành int, nếu lỗi trả về default
func atoi(s string, defaultVal int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultVal
	}
	return i
}

// parseBool chuyển string thành bool
func parseBool(s string, defaultVal bool) bool {
	if s == "" {
		return defaultVal
	}
	return strings.ToLower(s) == "true" || s == "1"
}

// GetConnections trả về tất cả database connections từ ENV
func GetConnections() map[string]ConnectionConfig {
	connections := make(map[string]ConnectionConfig)

	// SQLite connection
	// connections["sqlite"] = ConnectionConfig{
	// 	Driver:       "sqlite",
	// 	URL:          getEnv("DATABASE_URL", ""),
	// 	Database:     getEnv("DB_DATABASE", "database.sqlite"),
	// 	Prefix:       "",
	// 	ForeignKeys:  parseBool(getEnv("DB_FOREIGN_KEYS", "true"), true),
	// }

	// MySQL connection
	// connections["mysql"] = ConnectionConfig{
	// 	Driver:        "mysql",
	// 	URL:           getEnv("DATABASE_URL", ""),
	// 	Host:          getEnv("APP_DB_HOST", "127.0.0.1"),
	// 	Port:          getEnv("APP_DB_PORT", "3306"),
	// 	Database:      getEnv("APP_DB_DATABASE", "forge"),
	// 	Username:      getEnv("APP_DB_USERNAME", "forge"),
	// 	Password:      getEnv("APP_DB_PASSWORD", ""),
	// 	Charset:       "utf8mb4",
	// 	Collation:     "utf8mb4_unicode_ci",
	// 	Prefix:        "",
	// 	PrefixIndexes: true,
	// 	Strict:        true,
	// 	Engine:        "",
	// 	MaxOpenConns:  atoi(getEnv("DB_MAX_OPEN_CONNS", "10"), 10),
	// 	MaxIdleConns:  atoi(getEnv("DB_MAX_IDLE_CONNS", "5"), 5),
	// 	ConnMaxLifetime: parseDuration(getEnv("DB_CONN_MAX_LIFETIME", "30m")),
	// }

	// Data DB connection (dynamic)
	dataConnection := getEnv("DB_CONNECTION", "data")
	connections[dataConnection] = ConnectionConfig{
		Driver:          "mysql",
		URL:             getEnv("DATABASE_URL", ""),
		Host:            getEnv("DB_HOST", "127.0.0.1"),
		Port:            getEnv("DB_PORT", "3306"),
		Database:        getEnv("DB_DATABASE", "forge"),
		Username:        getEnv("DB_USERNAME", "forge"),
		Password:        getEnv("DB_PASSWORD", ""),
		Charset:         "utf8mb4",
		Collation:       "utf8mb4_unicode_ci",
		Prefix:          "",
		PrefixIndexes:   true,
		Strict:          true,
		ParseTime:       true,    // giá trị mặc định
		Loc:             "Local", // timezone mặc định
		Engine:          "",
		MaxOpenConns:    atoi(getEnv("DB_MAX_OPEN_CONNS", "10"), 10),
		MaxIdleConns:    atoi(getEnv("DB_MAX_IDLE_CONNS", "5"), 5),
		ConnMaxLifetime: parseDuration(getEnv("DB_CONN_MAX_LIFETIME", "30m")),
	}

	// App DB connection (dynamic)
	// appConnection := getEnv("APP_DB_CONNECTION", "app")
	// connections[appConnection] = ConnectionConfig{
	// 	Driver:        "mysql",
	// 	URL:           getEnv("DATABASE_URL", ""),
	// 	Host:          getEnv("APP_DB_HOST", "127.0.0.1"),
	// 	Port:          getEnv("APP_DB_PORT", "3306"),
	// 	Database:      getEnv("APP_DB_DATABASE", "forge"),
	// 	Username:      getEnv("APP_DB_USERNAME", "forge"),
	// 	Password:      getEnv("APP_DB_PASSWORD", ""),
	// 	Charset:       "utf8mb4",
	// 	Collation:     "utf8mb4_unicode_ci",
	// 	Prefix:        "",
	// 	PrefixIndexes: true,
	// 	Strict:        false,
	// 	Engine:        "",
	// 	MaxOpenConns:  atoi(getEnv("APP_DB_MAX_OPEN_CONNS", "10"), 10),
	// 	MaxIdleConns:  atoi(getEnv("APP_DB_MAX_IDLE_CONNS", "5"), 5),
	// 	ConnMaxLifetime: parseDuration(getEnv("APP_DB_CONN_MAX_LIFETIME", "30m")),
	// }

	// MySQL Logs connection
	// connections["mysql_logs"] = ConnectionConfig{
	// 	Driver:        "mysql",
	// 	URL:           getEnv("DATABASE_URL", ""),
	// 	Host:          getEnv("LOGS_DB_HOST", "127.0.0.1"),
	// 	Port:          getEnv("LOGS_DB_PORT", "3306"),
	// 	Database:      getEnv("LOGS_DB_DATABASE", "forge"),
	// 	Username:      getEnv("LOGS_DB_USERNAME", "forge"),
	// 	Password:      getEnv("LOGS_DB_PASSWORD", ""),
	// 	UnixSocket:    getEnv("DB_SOCKET", ""),
	// 	Charset:       "utf8mb4",
	// 	Collation:     "utf8mb4_unicode_ci",
	// 	Prefix:        "",
	// 	PrefixIndexes: true,
	// 	Strict:        true,
	// 	Engine:        "",
	// 	MaxOpenConns:  atoi(getEnv("LOGS_DB_MAX_OPEN_CONNS", "10"), 10),
	// 	MaxIdleConns:  atoi(getEnv("LOGS_DB_MAX_IDLE_CONNS", "5"), 5),
	// 	ConnMaxLifetime: parseDuration(getEnv("LOGS_DB_CONN_MAX_LIFETIME", "30m")),
	// }

	// PostgreSQL connection
	// connections["pgsql"] = ConnectionConfig{
	// 	Driver:        "postgres",
	// 	URL:           getEnv("DATABASE_URL", ""),
	// 	Host:          getEnv("DB_HOST", "127.0.0.1"),
	// 	Port:          getEnv("DB_PORT", "5432"),
	// 	Database:      getEnv("DB_DATABASE", "forge"),
	// 	Username:      getEnv("DB_USERNAME", "forge"),
	// 	Password:      getEnv("DB_PASSWORD", ""),
	// 	Charset:       "utf8",
	// 	Prefix:        "",
	// 	PrefixIndexes: true,
	// 	Schema:        "public",
	// 	SSLMode:       "prefer",
	// 	Timezone:      getEnv("DB_TIMEZONE", "Asia/Ho_Chi_Minh"),
	// 	MaxOpenConns:  atoi(getEnv("DB_MAX_OPEN_CONNS", "10"), 10),
	// 	MaxIdleConns:  atoi(getEnv("DB_MAX_IDLE_CONNS", "5"), 5),
	// 	ConnMaxLifetime: parseDuration(getEnv("DB_CONN_MAX_LIFETIME", "30m")),
	// }

	// SQL Server connection
	// connections["sqlsrv"] = ConnectionConfig{
	// 	Driver:        "sqlsrv",
	// 	URL:           getEnv("DATABASE_URL", ""),
	// 	Host:          getEnv("DB_HOST", "localhost"),
	// 	Port:          getEnv("DB_PORT", "1433"),
	// 	Database:      getEnv("DB_DATABASE", "forge"),
	// 	Username:      getEnv("DB_USERNAME", "forge"),
	// 	Password:      getEnv("DB_PASSWORD", ""),
	// 	Charset:       "utf8",
	// 	Prefix:        "",
	// 	PrefixIndexes: true,
	// 	MaxOpenConns:  atoi(getEnv("DB_MAX_OPEN_CONNS", "10"), 10),
	// 	MaxIdleConns:  atoi(getEnv("DB_MAX_IDLE_CONNS", "5"), 5),
	// 	ConnMaxLifetime: parseDuration(getEnv("DB_CONN_MAX_LIFETIME", "30m")),
	// }

	// MongoDB connection
	// connections["mongodb"] = ConnectionConfig{
	// 	Driver:   "mongodb",
	// 	URL:      getEnv("MONGODB_URI", ""),
	// 	Database: getEnv("MONGODB_DATABASE", ""),
	// 	MaxOpenConns:  atoi(getEnv("MONGODB_MAX_OPEN_CONNS", "10"), 10),
	// 	MaxIdleConns:  atoi(getEnv("MONGODB_MAX_IDLE_CONNS", "5"), 5),
	// 	ConnMaxLifetime: parseDuration(getEnv("MONGODB_CONN_MAX_LIFETIME", "30m")),
	// }

	return connections
}

// buildDSN xây dựng DSN từ connection config
func buildDSN(config ConnectionConfig) string {
	// Nếu có URL trực tiếp thì dùng luôn
	if config.URL != "" {
		return config.URL
	}

	switch config.Driver {
	case "mysql":
		return buildMySQLDSN(config)
	case "postgres":
		return buildPostgresDSN(config)
	case "sqlite":
		return buildSQLiteDSN(config)
	case "mongodb":
		return buildMongoDSN(config)
	case "sqlsrv":
		return buildSQLServerDSN(config)
	default:
		return ""
	}
}

// buildMySQLDSN xây dựng MySQL DSN
func buildMySQLDSN(config ConnectionConfig) string {
	host := config.Host
	if config.UnixSocket != "" {
		host = config.UnixSocket
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Username, config.Password, host, config.Port, config.Database,
		config.Charset, config.ParseTime, config.Loc)

	return dsn
}

// buildPostgresDSN xây dựng PostgreSQL DSN
func buildPostgresDSN(config ConnectionConfig) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.Host, config.Port, config.Username, config.Password, config.Database,
		config.SSLMode, config.Timezone)

	return dsn
}

// buildSQLiteDSN xây dựng SQLite DSN
func buildSQLiteDSN(config ConnectionConfig) string {
	return config.Database
}

// buildMongoDSN xây dựng MongoDB DSN
func buildMongoDSN(config ConnectionConfig) string {
	if config.URL != "" {
		return config.URL
	}

	// Xây dựng URI từ các thành phần
	if config.Username != "" && config.Password != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
			config.Username, config.Password, config.Host, config.Port, config.Database)
	}

	return fmt.Sprintf("mongodb://%s:%s/%s", config.Host, config.Port, config.Database)
}

// buildSQLServerDSN xây dựng SQL Server DSN
func buildSQLServerDSN(config ConnectionConfig) string {
	dsn := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s;charset=%s",
		config.Host, config.Port, config.Database, config.Username, config.Password, config.Charset)

	return dsn
}

// validateDBConfig kiểm tra cấu hình DB có hợp lệ không
func validateDBConfig(connectionName string, config ConnectionConfig) error {
	fmt.Println("Validating connection:", config)
	if config.Driver == "" {
		return fmt.Errorf("driver không được để trống cho connection: %s", connectionName)
	}

	// Kiểm tra driver được hỗ trợ
	supportedDrivers := []string{"mysql", "postgres", "sqlite", "mongodb", "sqlsrv"}
	isSupported := false
	for _, supported := range supportedDrivers {
		if config.Driver == supported {
			isSupported = true
			break
		}
	}
	if !isSupported {
		return fmt.Errorf("driver '%s' không được hỗ trợ cho connection: %s. Hỗ trợ: %v",
			config.Driver, connectionName, supportedDrivers)
	}

	return nil
}
