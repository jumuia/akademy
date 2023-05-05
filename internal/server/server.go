package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jumuia/akademy/internal/handler"
	"github.com/jumuia/akademy/internal/repository"
	"github.com/jumuia/akademy/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Port       string
	DBName     string
	DBUser     string
	DBPassword string
	DBURL      string
	DBHost     string
	DBPort     string
}

func Run(cfg *Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err = fmt.Errorf("unable to connect to the database: %w", err)
		return err
	}
	r := gin.Default()
	userRepo := repository.NewGORMUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	configureUserRoutes(r, userHandler)
	addr := fmt.Sprintf(":%s", cfg.Port)
	return r.Run(addr)
}

func configureUserRoutes(r *gin.Engine, h handler.UserHandler) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Pont @ %s", time.Now().String()),
		})
	})
	r.POST("/akademy/api/v1/users", h.CreateUser)
	r.GET("/akademy/api/v1/users", h.FetchAllUsers)
	r.GET("/akademy/api/v1/users/:user_id", h.FetchUser)
	r.PATCH("/akademy/api/v1/users/:user_id", h.UpdateUser)
	r.DELETE("/akademy/api/v1/users/:user_id", h.DeleteUser)
}
