package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thernande/app-pedidos-fondos-backend/internal/model"
	"github.com/thernande/app-pedidos-fondos-backend/internal/router/login_routes"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--migrate" {
		model.Migrate()
		os.Exit(0)
	}
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/", "./public")
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
	login := router.Group("/authentication")
	{
		login_routes.RoutesLogin(login)
	}
	router.Run(":8090")
}
