# Hướng dẫn sử dụng nhiều Database

## 🎯 **Cách lấy data từ các DB khác nhau**

### 1. **Lấy DB Instance**

```go
// Lấy các DB instances
mainDB := db.Get("main")           // DB chính
analyticsDB := db.Get("analytics") // DB analytics
logsDB := db.Get("logs")           // DB logs
```

### 2. **Lấy data từ DB chính**

```go
// Lấy tất cả users
var users []models.User
mainDB.Find(&users)

// Lấy user theo ID
var user models.User
mainDB.First(&user, userID)

// Lấy user với articles liên quan
var user models.User
mainDB.Preload("Articles").First(&user, userID)

// Tìm kiếm users
var users []models.User
mainDB.Where("username LIKE ?", "%john%").Find(&users)
```

### 3. **Lấy data từ DB analytics**

```go
// Lấy tất cả page views
var pageViews []models.PageView
analyticsDB.Find(&pageViews)

// Lấy page views theo ngày
var pageViews []models.PageView
analyticsDB.Where("DATE(visit_time) = ?", "2024-01-15").Find(&pageViews)

// Đếm page views
var count int64
analyticsDB.Model(&models.PageView{}).Count(&count)
```

### 4. **Lấy data từ DB logs**

```go
// Lấy logs theo level
var logs []models.Log
logsDB.Where("level = ?", "ERROR").Find(&logs)

// Lấy logs theo thời gian
var logs []models.Log
logsDB.Where("timestamp > ?", time.Now().Add(-24*time.Hour)).Find(&logs)
```

## 📝 **Ví dụ thực tế**

### **Tạo User và ghi Log**

```go
func CreateUserAndLog(username, email, password string) error {
    // 1. Tạo user trong DB chính
    mainDB := db.Get("main")
    user := &models.User{
        Username: username,
        Email:    email,
        Password: password,
    }
    
    if err := mainDB.Create(user).Error; err != nil {
        return err
    }
    
    // 2. Ghi log vào DB logs
    logsDB := db.Get("logs")
    log := &models.Log{
        Level:     "INFO",
        Message:   fmt.Sprintf("User %s created successfully", username),
        Timestamp: time.Now(),
    }
    
    if err := logsDB.Create(log).Error; err != nil {
        return err
    }
    
    return nil
}
```

### **Ghi Page View và Analytics**

```go
func RecordPageView(pageURL, userAgent, ipAddress string) error {
    // 1. Ghi page view vào DB analytics
    analyticsDB := db.Get("analytics")
    pageView := &models.PageView{
        PageURL:   pageURL,
        UserAgent: userAgent,
        IPAddress: ipAddress,
        VisitTime: time.Now(),
    }
    
    if err := analyticsDB.Create(pageView).Error; err != nil {
        return err
    }
    
    // 2. Ghi log vào DB logs
    logsDB := db.Get("logs")
    log := &models.Log{
        Level:     "INFO",
        Message:   fmt.Sprintf("Page view: %s from %s", pageURL, ipAddress),
        Timestamp: time.Now(),
    }
    
    if err := logsDB.Create(log).Error; err != nil {
        return err
    }
    
    return nil
}
```

### **Lấy Dashboard Data**

```go
func GetDashboardData() (map[string]interface{}, error) {
    result := make(map[string]interface{})
    
    // 1. Lấy data từ DB chính
    mainDB := db.Get("main")
    var userCount, articleCount int64
    mainDB.Model(&models.User{}).Count(&userCount)
    mainDB.Model(&models.Article{}).Count(&articleCount)
    
    result["users"] = userCount
    result["articles"] = articleCount
    
    // 2. Lấy data từ DB analytics
    analyticsDB := db.Get("analytics")
    var pageViewCount int64
    analyticsDB.Model(&models.PageView{}).Count(&pageViewCount)
    
    result["page_views"] = pageViewCount
    
    // 3. Lấy data từ DB logs
    logsDB := db.Get("logs")
    var errorLogCount int64
    logsDB.Model(&models.Log{}).Where("level = ?", "ERROR").Count(&errorLogCount)
    
    result["error_logs"] = errorLogCount
    
    return result, nil
}
```

## 🔄 **Transaction với nhiều DB**

### **Ví dụ: Tạo User với Log và Analytics**

```go
func CreateUserWithLogsAndAnalytics(user *models.User) error {
    // 1. Transaction trong DB chính
    mainDB := db.Get("main")
    tx := mainDB.Begin()
    
    if err := tx.Create(user).Error; err != nil {
        tx.Rollback()
        return err
    }
    
    // 2. Ghi log (không cần transaction vì logs không quan trọng)
    logsDB := db.Get("logs")
    log := &models.Log{
        Level:     "INFO",
        Message:   fmt.Sprintf("User %s created", user.Username),
        Timestamp: time.Now(),
    }
    logsDB.Create(log) // Bỏ qua lỗi vì logs không quan trọng
    
    // 3. Commit transaction chính
    if err := tx.Commit().Error; err != nil {
        return err
    }
    
    return nil
}
```

## 📊 **Kiểm tra trạng thái DB**

### **Kiểm tra DB có tồn tại không**

```go
// Cách an toàn - không panic nếu DB không tồn tại
if mainDB, ok := db.GetSafe("main"); ok {
    // DB tồn tại, sử dụng bình thường
    var users []models.User
    mainDB.Find(&users)
} else {
    // DB không tồn tại
    log.Println("DB main không tồn tại")
}
```

### **Xem danh sách tất cả DB**

```go
dbs := db.ListDBs()
for _, dbName := range dbs {
    log.Printf("DB: %s", dbName)
}
```

## 🎯 **Best Practices**

### 1. **Tách biệt logic theo DB**

```go
// Tốt: Tách riêng logic cho từng DB
func GetUserData(userID uint) (*models.User, error) {
    mainDB := db.Get("main")
    var user models.User
    return &user, mainDB.First(&user, userID).Error
}

func GetAnalyticsData() ([]models.PageView, error) {
    analyticsDB := db.Get("analytics")
    var pageViews []models.PageView
    return pageViews, analyticsDB.Find(&pageViews).Error
}
```

### 2. **Xử lý lỗi riêng biệt**

```go
func GetDataFromMultipleDBs() {
    // Lấy từ DB chính
    if users, err := GetUsersFromMainDB(); err != nil {
        log.Printf("Lỗi lấy users: %v", err)
    } else {
        log.Printf("Lấy được %d users", len(users))
    }
    
    // Lấy từ DB analytics
    if pageViews, err := GetPageViewsFromAnalytics(); err != nil {
        log.Printf("Lỗi lấy page views: %v", err)
    } else {
        log.Printf("Lấy được %d page views", len(pageViews))
    }
}
```

### 3. **Sử dụng constants cho DB names**

```go
const (
    DBMain      = "main"
    DBAnalytics = "analytics"
    DBLogs      = "logs"
)

func GetMainDB() *gorm.DB {
    return db.Get(DBMain)
}

func GetAnalyticsDB() *gorm.DB {
    return db.Get(DBAnalytics)
}
```

## 🚀 **Ví dụ hoàn chỉnh**

```go
func HandleUserRegistration(username, email, password string) error {
    // 1. Tạo user trong DB chính
    mainDB := db.Get("main")
    user := &models.User{
        Username: username,
        Email:    email,
        Password: password,
    }
    
    if err := mainDB.Create(user).Error; err != nil {
        return fmt.Errorf("lỗi tạo user: %v", err)
    }
    
    // 2. Ghi log
    logsDB := db.Get("logs")
    log := &models.Log{
        Level:     "INFO",
        Message:   fmt.Sprintf("User %s registered", username),
        Timestamp: time.Now(),
    }
    logsDB.Create(log)
    
    // 3. Ghi analytics
    analyticsDB := db.Get("analytics")
    registration := &models.UserActivity{
        UserID:    user.ID,
        Action:    "registration",
        Timestamp: time.Now(),
    }
    analyticsDB.Create(registration)
    
    return nil
}
```

Bây giờ bạn có thể dễ dàng lấy data từ nhiều database khác nhau! 🎉

