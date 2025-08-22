package database

import (
	"fmt"
	"log"
	"reflect"

	"gorm.io/gorm"
)

// ModelMapping chứa thông tin model và connection tương ứng
type ModelMapping struct {
	Model       interface{}
	Connection  string
	ModelName   string
}

// ModelRegistry chứa tất cả các model mapping cần migrate
var ModelRegistry = make([]ModelMapping, 0)

// RegisterModel đăng ký một model vào registry để migrate trên tất cả connections
func RegisterModel(model interface{}) {
	modelName := reflect.TypeOf(model).Name()
	ModelRegistry = append(ModelRegistry, ModelMapping{
		Model:      model,
		Connection: "*", // "*" nghĩa là tất cả connections
		ModelName:  modelName,
	})
	log.Printf("📝 Đã đăng ký model: %s (tất cả connections)", modelName)
}

// RegisterModelToConnection đăng ký một model vào registry để migrate trên connection cụ thể
func RegisterModelToConnection(model interface{}, connectionName string) {
	modelName := reflect.TypeOf(model).Name()
	ModelRegistry = append(ModelRegistry, ModelMapping{
		Model:      model,
		Connection: connectionName,
		ModelName:  modelName,
	})
	log.Printf("📝 Đã đăng ký model: %s (connection: %s)", modelName, connectionName)
}

// AutoMigrate thực hiện migrate tất cả các model đã đăng ký trên một connection
func AutoMigrate(db *gorm.DB) error {
	if len(ModelRegistry) == 0 {
		log.Println("⚠️ Không có model nào được đăng ký để migrate")
		return nil
	}

	log.Printf("🔄 Bắt đầu migrate %d models...", len(ModelRegistry))
	
	for _, mapping := range ModelRegistry {
		log.Printf("   🔄 Migrating model: %s", mapping.ModelName)
		
		if err := db.AutoMigrate(mapping.Model); err != nil {
			log.Printf("❌ Lỗi migrate model %s: %v", mapping.ModelName, err)
			return err
		}
		
		log.Printf("   ✅ Đã migrate thành công model: %s", mapping.ModelName)
	}
	
	log.Printf("🎉 Hoàn thành migrate %d models", len(ModelRegistry))
	return nil
}

// AutoMigrateAll thực hiện migrate trên tất cả các connections đã kết nối
func AutoMigrateAll() {
	log.Println("🚀 Bắt đầu migrate tất cả database connections...")
	
	for connectionName, db := range pool {
		log.Printf("📊 Migrating connection: %s", connectionName)
		
		// Lọc models cho connection này
		modelsForConnection := getModelsForConnection(connectionName)
		
		if len(modelsForConnection) == 0 {
			log.Printf("   ⚠️ Không có model nào cho connection: %s", connectionName)
			continue
		}
		
		log.Printf("   🔄 Migrating %d models cho connection: %s", len(modelsForConnection), connectionName)
		
		for _, mapping := range modelsForConnection {
			log.Printf("      🔄 Migrating model: %s", mapping.ModelName)
			
			if err := db.AutoMigrate(mapping.Model); err != nil {
				log.Printf("❌ Lỗi migrate model %s trên connection %s: %v", mapping.ModelName, connectionName, err)
			} else {
				log.Printf("      ✅ Đã migrate thành công model: %s trên connection: %s", mapping.ModelName, connectionName)
			}
		}
		
		log.Printf("✅ Hoàn thành migrate connection: %s", connectionName)
	}
}

// getModelsForConnection trả về danh sách models cho một connection cụ thể
func getModelsForConnection(connectionName string) []ModelMapping {
	var models []ModelMapping
	
	for _, mapping := range ModelRegistry {
		// Nếu model được đăng ký cho tất cả connections hoặc cho connection cụ thể này
		if mapping.Connection == "*" || mapping.Connection == connectionName {
			models = append(models, mapping)
		}
	}
	
	return models
}

// AutoMigrateConnection thực hiện migrate trên một connection cụ thể theo tên
func AutoMigrateConnection(connectionName string) error {
	db, ok := GetSafe(connectionName)
	if !ok {
		return fmt.Errorf("connection [%s] chưa được cấu hình", connectionName)
	}
	
	log.Printf("📊 Migrating connection: %s", connectionName)
	
	// Lọc models cho connection này
	modelsForConnection := getModelsForConnection(connectionName)
	
	if len(modelsForConnection) == 0 {
		log.Printf("⚠️ Không có model nào cho connection: %s", connectionName)
		return nil
	}
	
	log.Printf("🔄 Bắt đầu migrate %d models cho connection: %s", len(modelsForConnection), connectionName)
	
	for _, mapping := range modelsForConnection {
		log.Printf("   🔄 Migrating model: %s", mapping.ModelName)
		
		if err := db.AutoMigrate(mapping.Model); err != nil {
			log.Printf("❌ Lỗi migrate model %s: %v", mapping.ModelName, err)
			return err
		}
		
		log.Printf("   ✅ Đã migrate thành công model: %s", mapping.ModelName)
	}
	
	log.Printf("🎉 Hoàn thành migrate connection: %s", connectionName)
	return nil
}

// AutoMigrateDefaultConnection thực hiện migrate trên connection mặc định
func AutoMigrateDefaultConnection() error {
	defaultConnection := GetDefaultConnection()
	if defaultConnection == "" {
		return fmt.Errorf("không có connection mặc định nào được kết nối")
	}
	
	return AutoMigrateConnection(defaultConnection)
}

// ClearRegistry xóa tất cả các model đã đăng ký
func ClearRegistry() {
	ModelRegistry = ModelRegistry[:0]
	log.Println("🧹 Đã xóa tất cả models khỏi registry")
}

// ListRegisteredModels trả về danh sách tên các model đã đăng ký
func ListRegisteredModels() []string {
	models := make([]string, len(ModelRegistry))
	for i, mapping := range ModelRegistry {
		if mapping.Connection == "*" {
			models[i] = fmt.Sprintf("%s (tất cả connections)", mapping.ModelName)
		} else {
			models[i] = fmt.Sprintf("%s (connection: %s)", mapping.ModelName, mapping.Connection)
		}
	}
	return models
}

// GetModelMappings trả về chi tiết mapping của tất cả models
func GetModelMappings() []ModelMapping {
	return ModelRegistry
}

// GetModelsForConnection trả về danh sách models cho một connection cụ thể
func GetModelsForConnection(connectionName string) []ModelMapping {
	return getModelsForConnection(connectionName)
}

// GetModelsForAllConnections trả về danh sách models cho tất cả connections
func GetModelsForAllConnections() map[string][]ModelMapping {
	result := make(map[string][]ModelMapping)
	
	for connectionName := range pool {
		result[connectionName] = getModelsForConnection(connectionName)
	}
	
	return result
}

// RemoveModelFromRegistry xóa một model khỏi registry
func RemoveModelFromRegistry(modelName string) {
	for i, mapping := range ModelRegistry {
		if mapping.ModelName == modelName {
			ModelRegistry = append(ModelRegistry[:i], ModelRegistry[i+1:]...)
			log.Printf("🗑️ Đã xóa model %s khỏi registry", modelName)
			return
		}
	}
	log.Printf("⚠️ Không tìm thấy model %s trong registry", modelName)
}

// GetModelMappingByName trả về mapping của một model theo tên
func GetModelMappingByName(modelName string) (ModelMapping, bool) {
	for _, mapping := range ModelRegistry {
		if mapping.ModelName == modelName {
			return mapping, true
		}
	}
	return ModelMapping{}, false
}
