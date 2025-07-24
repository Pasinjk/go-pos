package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pasinjk/go-pos/internal/adapters/database"
	"github.com/pasinjk/go-pos/internal/adapters/http"
	"github.com/pasinjk/go-pos/internal/config"
	"github.com/pasinjk/go-pos/internal/domain/model"
	"github.com/pasinjk/go-pos/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	app := fiber.New()

	psql, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		})

	db, err := gorm.Open(postgres.Open(psql), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Category{},
		&model.Customer{},
		&model.Payment{},
		&model.Transaction{},
		&model.TransactionItem{},
		&model.StockMovement{},
	)

	userRepo := database.NewGormUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userHandler := http.NewHttpUserHandler(userService)

	app.Post("/register", userHandler.CreateUser)
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)

	categoriesRepo := database.NewGormCategoriesRepository(db)
	categoriesService := usecase.NewCategoriesService(categoriesRepo)
	categoriesHandler := http.NewHttpCategoriesHandler(categoriesService)

	app.Post("/categories", categoriesHandler.CreateCategory)
	app.Get("/categories", categoriesHandler.GetAllCategories)
	app.Get("/categories/:id", categoriesHandler.GetCategoryByID)
	app.Put("/categories/:id", categoriesHandler.UpdateCategory)
	app.Delete("/categories/:id", categoriesHandler.DeleteCatagoryByID)

	customerRepo := database.NewGormCustomerRepository(db)
	customerService := usecase.NewCustomerService(customerRepo)
	customerHandler := http.NewHttpCustomerHandler(customerService)

	app.Post("/customer", customerHandler.CreateCustomer)

	app.Listen(":8080")
}
