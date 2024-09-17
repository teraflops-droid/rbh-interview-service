package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/teraflops-droid/rbh-interview-service/configuration"
	"github.com/teraflops-droid/rbh-interview-service/controller"
	_ "github.com/teraflops-droid/rbh-interview-service/docs"
	"github.com/teraflops-droid/rbh-interview-service/exception"
	repository "github.com/teraflops-droid/rbh-interview-service/repository/impl"
	service "github.com/teraflops-droid/rbh-interview-service/service/impl"
)

//	@title			RBH-Interview APIs
//	@version		1.0
//	@description	Api documentation
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080
//	@BasePath		/
func main() {
	//setup configuration
	config := configuration.New()
	database := configuration.NewDatabase(config)
	//redis := configuration.NewRedis(config)

	//repository
	userRepository := repository.NewUserRepositoryImpl(database)

	//service
	userService := service.NewUserServiceImpl(&userRepository)

	//controller
	userController := controller.NewUserController(&userService, config)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())

	app.Use(recover.New())
	app.Use(cors.New())

	//swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//routing
	userController.Route(app)

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
