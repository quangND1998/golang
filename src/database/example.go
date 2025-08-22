package database

import (
	"fmt"
	"log"
)

// ExampleUsage demo cách sử dụng hệ thống database mới
func ExampleUsage() {
	log.Println("🚀 Demo Database System - Laravel Style")
	
	// 1. Khởi tạo database
	log.Println("\n1. Khởi tạo database connections...")
	Init()
	
	// 2. Liệt kê connections đã kết nối
	log.Println("\n2. Các connections đã kết nối:")
	connections := ListConnections()
	for _, conn := range connections {
		log.Printf("   ✅ %s", conn)
	}
	
	// 3. Lấy connection mặc định
	log.Println("\n3. Connection mặc định:")
	defaultConn := GetDefaultConnection()
	if defaultConn != "" {
		log.Printf("   🎯 Default: %s", defaultConn)
	} else {
		log.Println("   ⚠️ Không có connection mặc định")
	}
	
	// 4. Lấy thông tin chi tiết connections
	log.Println("\n4. Thông tin chi tiết connections:")
	allInfo := GetAllConnectionsInfo()
	for connName, info := range allInfo {
		log.Printf("   📊 %s:", connName)
		for key, value := range info {
			log.Printf("      %s: %v", key, value)
		}
	}
	
	// 5. Demo đăng ký models
	log.Println("\n5. Demo đăng ký models:")
	
	// Giả lập models
	type User struct {
		ID   uint   `gorm:"primarykey"`
		Name string `gorm:"size:255"`
		Email string `gorm:"size:255;unique"`
	}
	
	type Article struct {
		ID      uint   `gorm:"primarykey"`
		Title   string `gorm:"size:255"`
		Content string `gorm:"type:text"`
		UserID  uint
	}
	
	type Log struct {
		ID        uint   `gorm:"primarykey"`
		Level     string `gorm:"size:50"`
		Message   string `gorm:"type:text"`
		CreatedAt int64
	}
	
	// Đăng ký models
	RegisterModel(&User{})
	RegisterModel(&Article{})
	RegisterModelToConnection(&Log{}, "mysql_logs")
	
	// 6. Liệt kê models đã đăng ký
	log.Println("\n6. Models đã đăng ký:")
	registeredModels := ListRegisteredModels()
	for _, model := range registeredModels {
		log.Printf("   📝 %s", model)
	}
	
	// 7. Lấy models cho từng connection
	log.Println("\n7. Models cho từng connection:")
	modelsForConnections := GetModelsForAllConnections()
	for connName, models := range modelsForConnections {
		log.Printf("   🔗 %s:", connName)
		for _, model := range models {
			log.Printf("      - %s", model.ModelName)
		}
	}
	
	// 8. Demo migration (chỉ log, không thực hiện thực tế)
	log.Println("\n8. Demo migration:")
	for connName := range pool {
		log.Printf("   🔄 Sẽ migrate connection: %s", connName)
		models := GetModelsForConnection(connName)
		log.Printf("      Models: %d", len(models))
	}
	
	// 9. Demo quản lý connections
	log.Println("\n9. Demo quản lý connections:")
	
	// Kiểm tra connection có tồn tại không
	for _, connName := range connections {
		if IsConnected(connName) {
			log.Printf("   ✅ %s: Đã kết nối", connName)
		} else {
			log.Printf("   ❌ %s: Chưa kết nối", connName)
		}
	}
	
	// 10. Demo lấy connection an toàn
	log.Println("\n10. Demo lấy connection an toàn:")
	for _, connName := range []string{"mysql", "pgsql", "sqlite", "nonexistent"} {
		if db, exists := GetSafe(connName); exists {
			log.Printf("   ✅ %s: Kết nối thành công", connName)
			// Ping test
			sqlDB, err := db.DB()
			if err == nil {
				if err := sqlDB.Ping(); err == nil {
					log.Printf("      🏓 Ping: OK")
				} else {
					log.Printf("      ❌ Ping: %v", err)
				}
			}
		} else {
			log.Printf("   ❌ %s: Không tồn tại", connName)
		}
	}
	
	log.Println("\n🎉 Demo hoàn thành!")
}

// ExampleEnvFile tạo file .env mẫu
func ExampleEnvFile() {
	envContent := `# Database Configuration - Laravel Style

# Database URLs (optional)
DATABASE_URL=

# SQLite Configuration
DB_DATABASE=database.sqlite
DB_FOREIGN_KEYS=true

# MySQL Main Application Database
APP_DB_HOST=127.0.0.1
APP_DB_PORT=3306
APP_DB_DATABASE=myapp
APP_DB_USERNAME=root
APP_DB_PASSWORD=password
APP_DB_MAX_OPEN_CONNS=25
APP_DB_MAX_IDLE_CONNS=5
APP_DB_CONN_MAX_LIFETIME=5m

# Data Database (Dynamic Connection)
DATA_DB_CONNECTION=data
DATA_DB_HOST=127.0.0.1
DATA_DB_PORT=3306
DATA_DB_DATABASE=data_db
DATA_DB_USERNAME=data_user
DATA_DB_PASSWORD=data_pass
DATA_DB_MAX_OPEN_CONNS=10
DATA_DB_MAX_IDLE_CONNS=3
DATA_DB_CONN_MAX_LIFETIME=10m

# Logs Database
LOGS_DB_HOST=127.0.0.1
LOGS_DB_PORT=3306
LOGS_DB_DATABASE=logs_db
LOGS_DB_USERNAME=logs_user
LOGS_DB_PASSWORD=logs_pass
DB_SOCKET=
LOGS_DB_MAX_OPEN_CONNS=15
LOGS_DB_MAX_IDLE_CONNS=5
LOGS_DB_CONN_MAX_LIFETIME=15m

# PostgreSQL Database
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=postgres_db
DB_USERNAME=postgres
DB_PASSWORD=postgres_pass
DB_TIMEZONE=Asia/Ho_Chi_Minh
DB_MAX_OPEN_CONNS=20
DB_MAX_IDLE_CONNS=5
DB_CONN_MAX_LIFETIME=10m

# SQL Server Database
# DB_HOST=localhost
# DB_PORT=1433
# DB_DATABASE=sqlserver_db
# DB_USERNAME=sa
# DB_PASSWORD=password
# DB_MAX_OPEN_CONNS=10
# DB_MAX_IDLE_CONNS=3
# DB_CONN_MAX_LIFETIME=10m

# MongoDB Database
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=mongo_db
MONGODB_MAX_OPEN_CONNS=10
MONGODB_MAX_IDLE_CONNS=5
MONGODB_CONN_MAX_LIFETIME=30m
`
	
	fmt.Println("📄 File .env mẫu:")
	fmt.Println("```env")
	fmt.Println(envContent)
	fmt.Println("```")
}

// ExampleCode tạo code mẫu
func ExampleCode() {
	codeContent := `package main

import (
	"log"
	"your-project/internal/db"
	"your-project/internal/models"
)

func main() {
	// 1. Khởi tạo database
	db.Init()
	
	// 2. Đăng ký models
	db.RegisterModel(&models.User{})
	db.RegisterModel(&models.Article{})
	db.RegisterModelToConnection(&models.Log{}, "mysql_logs")
	db.RegisterModelToConnection(&models.Analytics{}, "data")
	
	// 3. Thực hiện migration
	db.AutoMigrateAll()
	
	// 4. Sử dụng databases
	// Lấy connection mặc định
	defaultDB := db.Get(db.GetDefaultConnection())
	
	// Lấy connection cụ thể
	mysqlDB := db.Get("mysql")
	pgsqlDB := db.Get("pgsql")
	logsDB := db.Get("mysql_logs")
	
	// 5. Thực hiện queries
	var users []models.User
	defaultDB.Find(&users)
	
	var logs []models.Log
	logsDB.Find(&logs)
	
	// 6. Kiểm tra thông tin connections
	connections := db.ListConnections()
	log.Printf("Connected databases: %v", connections)
	
	// 7. Lấy thông tin chi tiết
	info := db.GetConnectionInfo("mysql")
	log.Printf("MySQL info: %+v", info)
	
	// 8. Đóng connections khi kết thúc
	defer db.CloseAllConnections()
}
`
	
	fmt.Println("💻 Code mẫu:")
	fmt.Println("```go")
	fmt.Println(codeContent)
	fmt.Println("```")
}


