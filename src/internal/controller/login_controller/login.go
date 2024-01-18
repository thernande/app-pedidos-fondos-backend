package login_controller

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thernande/app-pedidos-fondos-backend/internal/config/database"
	loginhandler "github.com/thernande/app-pedidos-fondos-backend/internal/handler/login_handler"
	v1 "github.com/thernande/app-pedidos-fondos-backend/proto/login/v1"
	"google.golang.org/protobuf/proto"
)

func Login() {

}

func RegisterUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var req v1.RegisterUserRequest
	if err := proto.Unmarshal(body, &req); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := req.GetUser()
	handler := loginhandler.New(database.NewDb(database.New().MySQL()))
	handler.SetUserWithProtobuf(user)

	if err := handler.RegisterUser(); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := &v1.RegisterUserResponse{
		Success: true,
		Message: "Se ha registrado el usuario correctamente",
	}
	resBytes, _ := proto.Marshal(res)

	c.Data(http.StatusOK, "application/x-protobuf", resBytes)

}
