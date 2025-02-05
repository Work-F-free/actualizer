package handler

import (
	"actualizer/internal/handler/scheduler"
	"actualizer/internal/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

type Handler struct {
	schedulerHandler *scheduler.Handler
}

func New(srvs service.Scheduler) *Handler {
	return &Handler{
		schedulerHandler: scheduler.NewHandler(srvs),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n")
	}))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "refreshToken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("/actualizer")
	{
		service := api.Group("/service")
		{
			service.GET("/start", h.schedulerHandler.Start)
			service.GET("/stop", h.schedulerHandler.Stop)
		}

		booking := api.Group("/booking")
		{
			booking.GET("/cancel")
		}
	}

	return router
}
