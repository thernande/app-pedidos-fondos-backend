package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thernande/app-pedidos-fondos-backend/internal/controller/login_controller"
	"github.com/thernande/app-pedidos-fondos-backend/internal/model"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--migrate" {
		model.Migrate()
		os.Exit(0)
	}
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/", "./public")
	router.POST("/proto", login_controller.RegisterUser)
	router.Run(":8090")
}
