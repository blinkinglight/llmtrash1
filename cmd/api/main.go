package main

import (
	"fmt"
	"log"
	"net/http"

	"gotest/internal/adapter/driven/repository"
	adapter_http "gotest/internal/adapter/driving/http"
	"gotest/internal/domain/service"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 1. Initialize SQLite Connection (Pure Go)
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 2. Initialize Driven Adapter (SQLite Repository)
	userRepo := repository.NewSQLiteUserRepository(db)

	// 3. Initialize Domain Service (Use Case)
	userService := service.NewUserService(userRepo)

	// 4. Initialize Driving Adapter (HTTP Handler)
	userHandler := adapter_http.NewUserHandler(userService)

	// 5. Setup Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/users/register", userHandler.RegisterUser)
	mux.HandleFunc("/users/get", userHandler.GetUser)

	// 6. Start Server
	port := ":8080"
	fmt.Printf("Server starting on %s (SQLite mode)\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
