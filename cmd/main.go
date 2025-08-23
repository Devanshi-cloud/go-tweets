package main

import (
	"fmt"
	"go-tweets/internal/config"
	postHandler "go-tweets/internal/handler/post"
	userHandler "go-tweets/internal/handler/user"
	postRepo "go-tweets/internal/repository/post"
	userRepo "go-tweets/internal/repository/user"
	postService "go-tweets/internal/service/post"
	userService "go-tweets/internal/service/user"
	"go-tweets/pkg/internalsql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.Default()
	validate := validator.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db, err := internalsql.Connectmysql(cfg)
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/check-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "it works",
		})
	})

	userRepo := userRepo.NewRepository(db)
	postRepo := postRepo.NewPostRepository(db)

	userService := userService.NewService(cfg, userRepo)
	postService := postService.NewPostService(cfg, postRepo)

	userHandler := userHandler.NewHandler(r, validate, userService)
	postHandler := postHandler.NewHandler(r, validate, postService)

	postHandler.RouteList(cfg.SecretJwt)
	userHandler.RouteList(cfg.SecretJwt)

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)

	r.Run(server)
}