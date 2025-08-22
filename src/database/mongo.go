package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoPool = make(map[string]*mongo.Client)

// ConnectMongoFromEnv kết nối tất cả MongoDB từ ENV (đã được cập nhật cho cấu trúc mới)
func ConnectMongoFromEnv() {
	connections := GetConnections()

	log.Println("🔍 Tìm thấy MongoDB connections...")

	for connectionName, config := range connections {
		if config.Driver != "mongodb" {
			continue
		}

		log.Printf("🔗 Kết nối MongoDB [%s]...", connectionName)

		uri := buildMongoDSN(config)
		if uri == "" {
			log.Printf("⚠️ Bỏ qua MongoDB %s: không có URI", connectionName)
			continue
		}

		client, err := connectMongo(connectionName, uri, config.MaxOpenConns, config.MaxIdleConns, config.ConnMaxLifetime)
		if err != nil {
			log.Printf("❌ Lỗi kết nối MongoDB %s: %v", connectionName, err)
			continue
		}

		mongoPool[connectionName] = client
		log.Printf("✅ Kết nối MongoDB [%s] thành công", connectionName)
	}

	if len(mongoPool) == 0 {
		log.Println("⚠️ Không có MongoDB nào được kết nối thành công")
	} else {
		log.Printf("📊 Tổng cộng %d MongoDB đã được kết nối", len(mongoPool))
	}
}

// connectMongo kết nối MongoDB với connection config
func connectMongo(connectionName, uri string, maxOpenConns, maxIdleConns int, maxConnLifetime time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri).
		SetMaxPoolSize(uint64(maxOpenConns)).
		SetMinPoolSize(uint64(maxIdleConns)).
		SetMaxConnIdleTime(maxConnLifetime)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("không thể kết nối MongoDB: %v", err)
	}

	// Ping test
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping MongoDB error: %v", err)
	}

	return client, nil
}

// GetMongo trả về *mongo.Client theo tên connection
func GetMongo(connectionName string) *mongo.Client {
	client, ok := mongoPool[connectionName]
	if !ok {
		log.Fatalf("❌ MongoDB connection [%s] chưa được cấu hình", connectionName)
	}
	return client
}

// GetMongoSafe trả về *mongo.Client theo tên connection, không panic nếu không tìm thấy
func GetMongoSafe(connectionName string) (*mongo.Client, bool) {
	client, ok := mongoPool[connectionName]
	return client, ok
}

// GetMongoDB trả về *mongo.Database theo tên connection và database
func GetMongoDB(connectionName, databaseName string) *mongo.Database {
	client := GetMongo(connectionName)
	return client.Database(databaseName)
}

// ListMongoClients trả về danh sách tên các MongoDB connections đã kết nối
func ListMongoClients() []string {
	names := make([]string, 0, len(mongoPool))
	for name := range mongoPool {
		names = append(names, name)
	}
	return names
}

// CloseMongoClients đóng tất cả kết nối MongoDB
func CloseMongoClients() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for connectionName, client := range mongoPool {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("❌ Lỗi đóng MongoDB connection [%s]: %v", connectionName, err)
		} else {
			log.Printf("✅ Đã đóng MongoDB connection [%s]", connectionName)
		}
	}

	// Clear pool
	mongoPool = make(map[string]*mongo.Client)
}

// IsMongoConnected kiểm tra xem một MongoDB connection đã được kết nối chưa
func IsMongoConnected(connectionName string) bool {
	_, ok := mongoPool[connectionName]
	return ok
}

// GetMongoConnectionInfo trả về thông tin chi tiết của một MongoDB connection
func GetMongoConnectionInfo(connectionName string) map[string]interface{} {
	_, exists := GetMongoSafe(connectionName)
	if !exists {
		return nil
	}

	config, exists := GetConnectionConfig(connectionName)
	if !exists {
		return nil
	}

	info := map[string]interface{}{
		"driver":            config.Driver,
		"database":          config.Database,
		"uri":               maskSensitiveInfo(config.URL),
		"connected":         true,
		"max_open_conns":    config.MaxOpenConns,
		"max_idle_conns":    config.MaxIdleConns,
		"conn_max_lifetime": config.ConnMaxLifetime.String(),
	}

	return info
}

// GetAllMongoConnectionsInfo trả về thông tin tất cả MongoDB connections
func GetAllMongoConnectionsInfo() map[string]map[string]interface{} {
	info := make(map[string]map[string]interface{})

	for connectionName := range mongoPool {
		info[connectionName] = GetMongoConnectionInfo(connectionName)
	}

	return info
}

// CloseMongoConnection đóng một MongoDB connection cụ thể
func CloseMongoConnection(connectionName string) error {
	client, ok := mongoPool[connectionName]
	if !ok {
		return fmt.Errorf("MongoDB connection [%s] không tồn tại", connectionName)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		return fmt.Errorf("lỗi đóng MongoDB connection [%s]: %v", connectionName, err)
	}

	delete(mongoPool, connectionName)
	log.Printf("🔌 Đã đóng MongoDB connection [%s]", connectionName)
	return nil
}
