# Database Configuration - Laravel Style

Hệ thống database được thiết kế theo phong cách Laravel với multiple connections và cấu hình linh hoạt.

## Cấu trúc Connections

Hệ thống hỗ trợ các loại database connections sau:

### 1. SQLite
```env
DATABASE_URL=
DB_DATABASE=database.sqlite
DB_FOREIGN_KEYS=true
```

### 2. MySQL
```env
DATABASE_URL=
APP_DB_HOST=127.0.0.1
APP_DB_PORT=3306
APP_DB_DATABASE=forge
APP_DB_USERNAME=forge
APP_DB_PASSWORD=
APP_DB_MAX_OPEN_CONNS=10
APP_DB_MAX_IDLE_CONNS=5
APP_DB_CONN_MAX_LIFETIME=30m
```

### 3. Data DB (Dynamic)
```env
DATA_DB_CONNECTION=data
DATA_DB_HOST=127.0.0.1
DATA_DB_PORT=3306
DATA_DB_DATABASE=forge
DATA_DB_USERNAME=forge
DATA_DB_PASSWORD=
DATA_DB_MAX_OPEN_CONNS=10
DATA_DB_MAX_IDLE_CONNS=5
DATA_DB_CONN_MAX_LIFETIME=30m
```

### 4. App DB (Dynamic)
```env
APP_DB_CONNECTION=app
APP_DB_HOST=127.0.0.1
APP_DB_PORT=3306
APP_DB_DATABASE=forge
APP_DB_USERNAME=forge
APP_DB_PASSWORD=
APP_DB_MAX_OPEN_CONNS=10
APP_DB_MAX_IDLE_CONNS=5
APP_DB_CONN_MAX_LIFETIME=30m
```

### 5. MySQL Logs
```env
DATABASE_URL=
LOGS_DB_HOST=127.0.0.1
LOGS_DB_PORT=3306
LOGS_DB_DATABASE=forge
LOGS_DB_USERNAME=forge
LOGS_DB_PASSWORD=
DB_SOCKET=
LOGS_DB_MAX_OPEN_CONNS=10
LOGS_DB_MAX_IDLE_CONNS=5
LOGS_DB_CONN_MAX_LIFETIME=30m
```

### 6. PostgreSQL
```env
DATABASE_URL=
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=forge
DB_USERNAME=forge
DB_PASSWORD=
DB_TIMEZONE=Asia/Ho_Chi_Minh
DB_MAX_OPEN_CONNS=10
DB_MAX_IDLE_CONNS=5
DB_CONN_MAX_LIFETIME=30m
```

### 7. SQL Server
```env
DATABASE_URL=
DB_HOST=localhost
DB_PORT=1433
DB_DATABASE=forge
DB_USERNAME=forge
DB_PASSWORD=
DB_MAX_OPEN_CONNS=10
DB_MAX_IDLE_CONNS=5
DB_CONN_MAX_LIFETIME=30m
```

### 8. MongoDB
```env
MONGODB_URI=
MONGODB_DATABASE=
MONGODB_MAX_OPEN_CONNS=10
MONGODB_MAX_IDLE_CONNS=5
MONGODB_CONN_MAX_LIFETIME=30m
```

## Sử dụng trong Code

### Khởi tạo Database
```go
package main

import "your-project/internal/db"

func main() {
    // Khởi tạo tất cả database connections
    db.Init()
}
```

### Lấy Database Connection
```go
// Lấy connection mặc định
defaultDB := db.Get(db.GetDefaultConnection())

// Lấy connection cụ thể
mysqlDB := db.Get("mysql")
pgsqlDB := db.Get("pgsql")
sqliteDB := db.Get("sqlite")

// Lấy connection an toàn (không panic)
if db, exists := db.GetSafe("mysql"); exists {
    // Sử dụng db
}
```

### Đăng ký Models cho Migration
```go
import "your-project/internal/models"

// Đăng ký model cho tất cả connections
db.RegisterModel(&models.User{})
db.RegisterModel(&models.Article{})

// Đăng ký model cho connection cụ thể
db.RegisterModelToConnection(&models.Log{}, "mysql_logs")
db.RegisterModelToConnection(&models.Analytics{}, "data")
```

### Thực hiện Migration
```go
// Migrate tất cả connections
db.AutoMigrateAll()

// Migrate connection cụ thể
db.AutoMigrateConnection("mysql")

// Migrate connection mặc định
db.AutoMigrateDefaultConnection()

// Migrate trực tiếp trên DB instance
db.AutoMigrate(db.Get("mysql"))
```

### Quản lý Connections
```go
// Liệt kê tất cả connections đã kết nối
connections := db.ListConnections()
fmt.Println("Connected:", connections)

// Kiểm tra connection có tồn tại không
if db.IsConnected("mysql") {
    fmt.Println("MySQL đã được kết nối")
}

// Lấy thông tin connection
info := db.GetConnectionInfo("mysql")
fmt.Printf("MySQL Info: %+v\n", info)

// Lấy thông tin tất cả connections
allInfo := db.GetAllConnectionsInfo()
fmt.Printf("All Connections: %+v\n", allInfo)

// Đóng connection cụ thể
db.CloseConnection("mysql")

// Đóng tất cả connections
db.CloseAllConnections()
```

### Quản lý Models
```go
// Liệt kê models đã đăng ký
models := db.ListRegisteredModels()
fmt.Println("Registered Models:", models)

// Lấy models cho connection cụ thể
modelsForMySQL := db.GetModelsForConnection("mysql")

// Lấy models cho tất cả connections
allModels := db.GetModelsForAllConnections()

// Xóa model khỏi registry
db.RemoveModelFromRegistry("User")

// Lấy mapping của model
if mapping, exists := db.GetModelMappingByName("User"); exists {
    fmt.Printf("User mapping: %+v\n", mapping)
}
```

## Ví dụ File .env

```env
# Database URLs
DATABASE_URL=

# SQLite
DB_DATABASE=database.sqlite
DB_FOREIGN_KEYS=true

# MySQL Main
APP_DB_HOST=127.0.0.1
APP_DB_PORT=3306
APP_DB_DATABASE=myapp
APP_DB_USERNAME=root
APP_DB_PASSWORD=password
APP_DB_MAX_OPEN_CONNS=25
APP_DB_MAX_IDLE_CONNS=5
APP_DB_CONN_MAX_LIFETIME=5m

# Data Database
DATA_DB_CONNECTION=data
DATA_DB_HOST=127.0.0.1
DATA_DB_PORT=3306
DATA_DB_DATABASE=data_db
DATA_DB_USERNAME=data_user
DATA_DB_PASSWORD=data_pass

# Logs Database
LOGS_DB_HOST=127.0.0.1
LOGS_DB_PORT=3306
LOGS_DB_DATABASE=logs_db
LOGS_DB_USERNAME=logs_user
LOGS_DB_PASSWORD=logs_pass

# PostgreSQL
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=postgres_db
DB_USERNAME=postgres
DB_PASSWORD=postgres_pass
DB_TIMEZONE=Asia/Ho_Chi_Minh

# MongoDB
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=mongo_db
```

## Lưu ý

1. **MongoDB**: Hiện tại chưa được implement đầy đủ, chỉ log thông tin
2. **Connection Pool**: Mỗi connection có thể cấu hình pool riêng
3. **Migration**: Models có thể được đăng ký cho tất cả connections hoặc connection cụ thể
4. **Validation**: Hệ thống tự động validate cấu hình trước khi kết nối
5. **Error Handling**: Các lỗi kết nối được log chi tiết và không làm crash ứng dụng
