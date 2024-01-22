package login_routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thernande/app-pedidos-fondos-backend/internal/controller/login_controller"
)

// RoutesLogin : configures routes for login controller
func RoutesLogin(router *gin.RouterGroup) {
	router.POST("/loginUser", login_controller.Login)
	router.PUT("/registerUser", login_controller.RegisterUser)
}
