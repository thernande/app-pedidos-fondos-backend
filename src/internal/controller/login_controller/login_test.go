package login_controller

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	v1 "github.com/thernande/app-pedidos-fondos-backend/proto/login/v1"
	"google.golang.org/protobuf/proto"
)

func TestRegisterUser(t *testing.T) {
	// Initialize test data with random values
	rand.Seed(time.Now().UnixNano())
	//document := strconv.Itoa(rand.Intn(1000000000) + 1000000000) // Random 10-digit number
	document := "1234567890"
	password := "testPassword" + strconv.Itoa(rand.Intn(100))
	name := "John" + strconv.Itoa(rand.Intn(100))
	lastname := "Doe" + strconv.Itoa(rand.Intn(100))
	email := fmt.Sprintf("john.doe%d@example.com", rand.Intn(100))
	phone := fmt.Sprintf("%03d-%03d-%04d", rand.Intn(1000), rand.Intn(1000), rand.Intn(10000))
	company := "TestCorp" + strconv.Itoa(rand.Intn(100))

	requestData := &v1.RegisterUserRequest{
		User: &v1.User{
			Document: document,
			Password: password,
			Name:     name,
			Lastname: lastname,
			Email:    email,
			Phone:    phone,
			Company:  company,
		},
	}
	requestDataBytes, _ := proto.Marshal(requestData)
	request := httptest.NewRequest(http.MethodPost, "/proto", bytes.NewReader(requestDataBytes))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = request

	// Invoke the function
	RegisterUser(c)

	// Assert the response
	response := &v1.RegisterUserResponse{}
	proto.Unmarshal(w.Body.Bytes(), response)
	fmt.Println(response)
	assert.True(t, response.Success)
	assert.Equal(t, "Se ha registrado el usuario correctamente", response.Message)
	if !assert.True(t, response.Success) {
		t.Logf("TestRegisterUser failed, expected response.Success to be true, got false")
	}
	if !assert.Equal(t, "Se ha registrado el usuario correctamente", response.Message) {
		t.Logf("TestRegisterUser failed, expected message to be 'Se ha registrado el usuario correctamente', got '%s'", response.Message)
	}
}
