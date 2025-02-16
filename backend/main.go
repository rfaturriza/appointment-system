package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rfaturriza/appointment-system/config"
	"github.com/rfaturriza/appointment-system/db"
	"github.com/rfaturriza/appointment-system/handlers"
	"github.com/rfaturriza/appointment-system/middleware"
)

func main() {
    config := config.LoadConfig()
    db.InitDB()

    router := gin.Default()

    corsConfig := cors.DefaultConfig()
    corsConfig.AllowAllOrigins = true
    corsConfig.AllowCredentials = true
    corsConfig.AddAllowHeaders("Authorization")
    corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
    router.Use(cors.New(corsConfig))

    // Trust only local proxies
    router.SetTrustedProxies([]string{"127.0.0.1"})
    router.GET("/ping", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "message": "pong",
      })
    })
    router.POST("/login", handlers.Login)
    router.Use(middleware.AuthMiddleware())
    router.GET("/users/:id", handlers.GetUser)
    router.GET("/users/me", handlers.GetCurrentUser)
    router.POST("/appointments", handlers.CreateAppointment)
    router.GET("/appointments", handlers.GetUpcomingAppointments)


    router.Run(config.Port)
}