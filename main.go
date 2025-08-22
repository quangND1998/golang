package main

import (
	"context"
	"log"
	"nextlend-api-web-frontend/src/common/response"
	"nextlend-api-web-frontend/src/configuration"
	controller "nextlend-api-web-frontend/src/controllers"
	"nextlend-api-web-frontend/src/database"
	"nextlend-api-web-frontend/src/exception"
	"nextlend-api-web-frontend/src/middleware"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	repository "nextlend-api-web-frontend/src/repository/impl"
	service "nextlend-api-web-frontend/src/service/impl"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	config := configuration.New()
	database.Init()

	app := fiber.New(configuration.NewFiberConfiguration())
	newsCategoryRepository := repository.NewsCategoryRepositoryImpl()
	mewCategoryService := service.NewCategoryServiceImpl(&newsCategoryRepository)

	newCategoryController := controller.InitNewCategoryController(&mewCategoryService)
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(middleware.LoggingMiddleware())
	app.Get("/healthz", func(c *fiber.Ctx) error {
		sqlConnections := database.ListConnections()
		mongoConnections := database.ListMongoClients()
		return response.Success(c, fiber.Map{
			"status":  "ok",
			"service": "fiber-starter",
			"time":    time.Now(),
			"databases": fiber.Map{
				"sql":     sqlConnections,
				"mongodb": mongoConnections,
			},
		})

	})
	//swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	newCategoryController.Route(app)
	//start app
	err := app.Listen(":" + config.Get("PORT"))
	exception.PanicLogging(err)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Đang tắt server...")

	// Đóng database connections
	log.Println("🔌 Đang đóng database connections...")
	database.CloseAllConnections()
	database.CloseMongoClients()

	// Shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("❌ Lỗi khi tắt server: %v", err)
	}

	log.Println("✅ Server đã tắt thành công")

}
