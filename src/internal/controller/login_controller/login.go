package login_controller

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thernande/app-pedidos-fondos-backend/internal/config/database"
	loginhandler "github.com/thernande/app-pedidos-fondos-backend/internal/handler/login_handler"
	v1 "github.com/thernande/app-pedidos-fondos-backend/proto/login/v1"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
)

func Login(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var req v1.LoginRequest
	if err := proto.Unmarshal(body, &req); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	handler := loginhandler.NewLoginUser(database.NewDb(database.New().MySQL()))
	handler.User.Documento = req.Document
	if err := handler.LoginUser(req.Password); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := &v1.LoginResponse{
		Token: handler.Token,
	}
	resBytes, err := proto.Marshal(res)
	if err != nil {
		handler.Db.Log.ErrorLogPrint(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Data(http.StatusOK, "application/x-protobuf", resBytes)
}

/*
	RegisterUser handles the registration of a user.

This function reads the request body and unmarshals it into a RegisterUserRequest protobuf message.
If there is an error reading the request body or unmarshaling the message, it returns a RegisterUserResponse protobuf message with the error details.

It then validates the user's password length and returns an error if it is less than 8 characters or more than 30 characters.

If the password is valid, it generates a hashed password using bcrypt and updates the user's password field.

It creates a new instance of the RegisterUser handler and sets the user with the protobuf message.
It then calls the RegisterUser method of the handler to register the user.
If there is an error registering the user, it returns a RegisterUserResponse protobuf message with the error details.

If the user is successfully registered, it returns a RegisterUserResponse protobuf message with a success status and a success message.

Parameters:
- v1.RegisterUserRequest: The request body of the registration request.
*/
func RegisterUser(c *gin.Context) {
	fmt.Println("RegisterUser")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		res := &v1.RegisterUserResponse{
			Success: false,
			Message: err.Error(),
		}
		resBytes, _ := proto.Marshal(res)
		c.Data(http.StatusBadRequest, "application/x-protobuf", resBytes)
		return
	}
	var req v1.RegisterUserRequest
	if err := proto.Unmarshal(body, &req); err != nil {
		fmt.Println(err)
		res := &v1.RegisterUserResponse{
			Success: false,
			Message: err.Error(),
		}
		resBytes, _ := proto.Marshal(res)
		c.Data(http.StatusBadRequest, "application/x-protobuf", resBytes)
		return
	}

	user := req.GetUser()
	if len(user.Password) < 8 && len(user.Password) > 30 {
		res := &v1.RegisterUserResponse{
			Success: false,
			Message: "La contraseña debe tener al menos 8 caracteres y no más de 30 caracteres",
		}
		resBytes, _ := proto.Marshal(res)
		c.Data(http.StatusBadRequest, "application/x-protobuf", resBytes)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		res := &v1.RegisterUserResponse{
			Success: false,
			Message: err.Error(),
		}
		resBytes, _ := proto.Marshal(res)
		c.Data(http.StatusInternalServerError, "application/x-protobuf", resBytes)
		return
	}
	user.Password = string(hashedPassword)

	handler := loginhandler.NewRegisterUser(database.NewDb(database.New().MySQL()))
	handler.SetUserWithProtobuf(user)
	if err := handler.RegisterUser(); err != nil {
		res := &v1.RegisterUserResponse{
			Success: false,
			Message: err.Error(),
		}
		resBytes, _ := proto.Marshal(res)
		c.Data(http.StatusBadRequest, "application/x-protobuf", resBytes)
		return
	}

	res := &v1.RegisterUserResponse{
		Success: true,
		Message: "Se ha registrado el usuario correctamente",
	}
	resBytes, _ := proto.Marshal(res)

	c.Data(http.StatusOK, "application/x-protobuf", resBytes)
}

// This function encrypts the provided password 20 times using bcrypt and returns the final hashed password.
func EncryptPasswordMultipleTimes() {
	password := "testPassword"
	var hashedPassword []byte
	var err error

	for i := 0; i < 20; i++ {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		fmt.Println(string(hashedPassword))
	}
}
