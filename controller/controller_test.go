package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/teraflops-droid/rbh-interview-service/configuration"
	"github.com/teraflops-droid/rbh-interview-service/exception"
	"github.com/teraflops-droid/rbh-interview-service/model"
	"github.com/teraflops-droid/rbh-interview-service/repository/impl"
	impl2 "github.com/teraflops-droid/rbh-interview-service/service/impl"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http/httptest"
)

func createTestApp() *fiber.App {
	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	//routing
	userController.Route(app)

	return app
}

// setup configuration
var config = configuration.New("../.env.test")
var database = configuration.NewDatabase(config)
var redis = configuration.NewRedis(config)

// repository
var userRepository = impl.NewUserRepositoryImpl(database)

// service
var userService = impl2.NewUserServiceImpl(&userRepository)

// controller
var userController = NewUserController(&userService, config)

var appTest = createTestApp()

func authenticationCreate() map[string]interface{} {
	userRepository.DeleteAll()

	password, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	exception.PanicLogging(err)
	roles := []string{"ROLE_ADMIN", "ROLE_USER"}
	userRepository.Create("admin", string(password), roles)

	userModel := model.UserModel{
		Username: "admin",
		Password: "admin",
	}

	userRequestBody, _ := json.Marshal(userModel)

	userRequest := httptest.NewRequest("POST", "/v1/api/authentication", bytes.NewBuffer(userRequestBody))
	userRequest.Header.Set("Content-Type", "application/json")
	userRequest.Header.Set("Accept", "application/json")

	userResponse, _ := appTest.Test(userRequest)

	userResponseBody, _ := io.ReadAll(userResponse.Body)
	userWebResponse := model.GeneralResponse{}
	_ = json.Unmarshal(userResponseBody, &userWebResponse)

	userJsonData, _ := json.Marshal(userWebResponse.Data)

	tokenResponse := map[string]interface{}{}
	_ = json.Unmarshal(userJsonData, &tokenResponse)

	return tokenResponse
}
