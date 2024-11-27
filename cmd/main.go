package main

import (
	"Ecommerce/config"
	"Ecommerce/internal/handler"
	"Ecommerce/internal/repository"
	"Ecommerce/internal/service"
	"Ecommerce/pkg/database"
	"Ecommerce/pkg/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	// 公開路由
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// 需要認證的路由組
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{

	}
	log.Fatal(r.Run(":8080"))
}
