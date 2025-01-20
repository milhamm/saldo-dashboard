package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"saldo-server/config"
	"saldo-server/handler"
	"saldo-server/repository"
	"saldo-server/service"
	"time"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pool := config.InitDB()
	defer pool.Close()

	userRepository := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepository)

	movementRepository := repository.NewMovementRepository(pool)
	movementService := service.NewMovementService(movementRepository, userRepository)

	accountRepository := repository.NewAccountRepository(pool)
	accountService := service.NewAccountService(accountRepository)

	authService := service.NewAuthService(userRepository)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	authMiddlerware := echojwt.WithConfig(config.InitJWT())

	publicApiGroup := e.Group("/api/v1")
	privateApiGroup := e.Group("/api/v1", authMiddlerware)

	userRouteGroup := privateApiGroup.Group("/users")
	handler.NewUserHandler(userRouteGroup, userService)

	movementRouteGroup := privateApiGroup.Group("/movements")
	handler.NewMovementHandler(movementRouteGroup, movementService)

	accountRouteGroup := privateApiGroup.Group("/accounts")
	handler.NewAccountHandler(accountRouteGroup, accountService)

	authRouteGroup := publicApiGroup.Group("/auth")
	handler.NewAuthHandler(authRouteGroup, authService)
	startServer(e)
}

func startServer(e *echo.Echo) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":5000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("shutting down the server %v\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// $ migrate create -ext sql -dir ./migration/ -seq init_mg
// $ migrate -database "postgres://postgres:postgres@db:5432/saldo-db?sslmode=disable" -path ./apps/server/migrations up
