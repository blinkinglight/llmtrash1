package main

import (
	"fmt"
	"net/http"

	"gotest/internal/adapter/driven/repository"
	webapi "gotest/internal/adapter/driving/http"
	"gotest/internal/domain/service"
)

func main() {
	// 1. Initialize Driven Adapter (Repository)
	userRepo := repository.NewInMemoryUserRepository()

	// 2. Initialize Domain Service (Use Case)
	userService := service.NewUserService(userRepo)

	// 3. Initialize Driving Adapter (HTTP Handler)
	webHandler := webapi.NewUserHandler(userService)

	// 4. Setup Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/users/register", webHandler.RegisterUser)
	mux.HandleFunc("/users/get", webHandler.GetUser)

	// 5. Start Server
	port := ":8080"
	fmt.Printf("Server starting on %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
