# Hệ thống Auto Migrate linh hoạt với Database Mapping

Hệ thống này cho phép bạn dễ dàng quản lý việc migrate các models và chỉ định chúng sẽ được migrate ở database nào.

## Cách sử dụng

### 1. Đăng ký Model

#### Đăng ký cho tất cả DB (mặc định)
```go
// Trong main.go
func setupAndMigrateModels() {
    // Models cho tất cả DB
    db.RegisterModel(&models.User{})
    db.RegisterModel(&models.Follow{})
    db.RegisterModel(&models.Article{})
    db.RegisterModel(&models.Comment{})
    db.RegisterModel(&models.Tag{})
    
    // Thêm model mới cho tất cả DB
    db.RegisterModel(&models.YourNewModel{})
    
    db.AutoMigrateAll()
}
```

#### Đăng ký cho DB cụ thể
```go
func setupAndMigrateModels() {
    // Models cho DB cụ thể
    db.RegisterModelToDB(&models.User{}, "main")
    db.RegisterModelToDB(&models.Product{}, "analytics")
    db.RegisterModelToDB(&models.Log{}, "logs")
    
    // Models cho tất cả DB
    db.RegisterModel(&models.Article{})
    
    db.AutoMigrateAll()
}
```

### 2. Các hàm có sẵn

#### `RegisterModel(model interface{})`
Đăng ký một model vào registry để migrate trên tất cả DB.

```go
db.RegisterModel(&models.User{})
```

#### `RegisterModelToDB(model interface{}, dbName string)`
Đăng ký một model vào registry để migrate trên DB cụ thể.

```go
db.RegisterModelToDB(&models.User{}, "main")
db.RegisterModelToDB(&models.Product{}, "analytics")
```

#### `AutoMigrate(db *gorm.DB) error`
Migrate tất cả models đã đăng ký trên một DB cụ thể.

```go
db := db.Get("main")
err := db.AutoMigrate(db)
```

#### `AutoMigrateAll()`
Migrate tất cả models trên tất cả databases đã kết nối (theo mapping).

```go
db.AutoMigrateAll()
```

#### `AutoMigrateDB(dbName string) error`
Migrate tất cả models trên một DB cụ thể theo tên.

```go
err := db.AutoMigrateDB("main")
```

#### `ListRegisteredModels() []string`
Xem danh sách tất cả models đã đăng ký với thông tin DB.

```go
models := db.ListRegisteredModels()
for _, model := range models {
    fmt.Println("Model:", model)
}
```

#### `GetModelMappings() []ModelMapping`
Xem chi tiết mapping của tất cả models.

```go
mappings := db.GetModelMappings()
for _, mapping := range mappings {
    fmt.Printf("Model: %s, DB: %s\n", mapping.ModelName, mapping.DBName)
}
```

#### `GetModelsForDatabase(dbName string) []ModelMapping`
Xem danh sách models cho một database cụ thể.

```go
models := db.GetModelsForDatabase("main")
for _, mapping := range models {
    fmt.Printf("Model: %s\n", mapping.ModelName)
}
```

#### `ClearRegistry()`
Xóa tất cả models khỏi registry.

```go
db.ClearRegistry()
```

### 3. Ví dụ sử dụng

#### Thêm model mới cho tất cả DB

1. Tạo model trong `internal/models/`:

```go
// internal/models/product.go
package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name        string  `gorm:"not null"`
    Price       float64 `gorm:"not null"`
    Description string
}
```

2. Đăng ký model trong `main.go`:

```go
func setupAndMigrateModels() {
    // ... models hiện có ...
    
    // Thêm model mới cho tất cả DB
    db.RegisterModel(&models.Product{})
    
    db.AutoMigrateAll()
}
```

3. Chạy ứng dụng - model sẽ được migrate tự động trên tất cả DB!

#### Thêm model mới cho DB cụ thể

1. Tạo model cho analytics:

```go
// internal/models/analytics.go
package models

import "gorm.io/gorm"

type PageView struct {
    gorm.Model
    PageURL     string    `gorm:"not null"`
    UserAgent   string
    IPAddress   string
    VisitTime   time.Time `gorm:"not null"`
}
```

2. Đăng ký model cho DB analytics:

```go
func setupAndMigrateModels() {
    // Models cho DB chính
    db.RegisterModelToDB(&models.User{}, "main")
    db.RegisterModelToDB(&models.Article{}, "main")
    
    // Models cho DB analytics
    db.RegisterModelToDB(&models.PageView{}, "analytics")
    
    db.AutoMigrateAll()
}
```

3. Model PageView sẽ chỉ được migrate trên DB "analytics"!

#### Migrate một DB cụ thể

```go
// Migrate chỉ DB "main"
err := db.AutoMigrateDB("main")
if err != nil {
    log.Printf("Lỗi migrate: %v", err)
}
```

#### Xem trạng thái

```go
// Xem danh sách DB đã kết nối
dbs := db.ListDBs()
fmt.Println("Databases:", dbs)

// Xem danh sách models đã đăng ký với thông tin DB
models := db.ListRegisteredModels()
for _, model := range models {
    fmt.Println("Model:", model)
}

// Xem chi tiết mapping
mappings := db.GetModelMappings()
for _, mapping := range mappings {
    fmt.Printf("Model: %s → DB: %s\n", mapping.ModelName, mapping.DBName)
}

// Xem models cho DB cụ thể
mainModels := db.GetModelsForDatabase("main")
fmt.Printf("Models cho DB 'main': %d\n", len(mainModels))
for _, mapping := range mainModels {
    fmt.Printf("  - %s\n", mapping.ModelName)
}
```

## Lợi ích

1. **Đơn giản**: Chỉ cần thêm 1 dòng để đăng ký model mới
2. **Linh hoạt**: Có thể chỉ định model migrate ở DB nào
3. **Tự động**: Migrate tự động khi khởi động ứng dụng
4. **Rõ ràng**: Biết chính xác model nào ở DB nào
5. **An toàn**: Có logging chi tiết và xử lý lỗi
6. **Dễ quản lý**: Có thể xem danh sách models và DBs

## Log Output

Khi chạy migrate, bạn sẽ thấy output như sau:

```
📝 Đã đăng ký model: User (tất cả DB)
📝 Đã đăng ký model: Follow (tất cả DB)
📝 Đã đăng ký model: Article (tất cả DB)
📝 Đã đăng ký model: Comment (tất cả DB)
📝 Đã đăng ký model: Tag (tất cả DB)
📝 Đã đăng ký model: PageView (DB: analytics)
🚀 Bắt đầu migrate tất cả databases...
📊 Migrating database: main
   🔄 Migrating 5 models cho DB: main
      🔄 Migrating model: User
      ✅ Đã migrate thành công model: User trên DB: main
      🔄 Migrating model: Follow
      ✅ Đã migrate thành công model: Follow trên DB: main
      ...
✅ Hoàn thành migrate DB: main
📊 Migrating database: analytics
   🔄 Migrating 6 models cho DB: analytics
      🔄 Migrating model: User
      ✅ Đã migrate thành công model: User trên DB: analytics
      ...
      🔄 Migrating model: PageView
      ✅ Đã migrate thành công model: PageView trên DB: analytics
✅ Hoàn thành migrate DB: analytics
📊 Thông tin model mapping:
   User → tất cả DB
   Follow → tất cả DB
   Article → tất cả DB
   Comment → tất cả DB
   Tag → tất cả DB
   PageView → DB: analytics
```
