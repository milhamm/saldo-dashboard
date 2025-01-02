package main

import (
	"log"
	"os"
	"saldo-server/config"
	"saldo-server/handler"
	"saldo-server/repository"
	"saldo-server/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pool := config.InitDB(os.Getenv("DATABASE_URL"))
	defer pool.Close()

	userRepository := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepository)

	e := echo.New()
	e.Use(middleware.Logger())

	apiGroup := e.Group("/api/v1")

	userRouteGroup := apiGroup.Group("/users")
	handler.NewUserHandler(userRouteGroup, userService)

	e.Logger.Fatal(e.Start(":5000"))
}

// $ migrate create -ext sql -dir ./migration/ -seq init_mg
// $ migrate -database "postgres://postgres:postgres@db:5432/saldo-db?sslmode=disable" -path ./apps/server/migrations up
