package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/teraflops-droid/rbh-interview-service/common/logger"
	"github.com/teraflops-droid/rbh-interview-service/configuration"
	"github.com/teraflops-droid/rbh-interview-service/controller"
	_ "github.com/teraflops-droid/rbh-interview-service/docs"
	"github.com/teraflops-droid/rbh-interview-service/exception"
	repository "github.com/teraflops-droid/rbh-interview-service/repository/impl"
	service "github.com/teraflops-droid/rbh-interview-service/service/impl"
)

// @title			RBH-Interview APIs
// @version		1.0
// @description	Api documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	Nadthapon Sukeewadthana
// @contact.email	nadthapon1998@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
func main() {
	//setup configuration
	config := configuration.New()
	database := configuration.NewDatabase(config)
	//redis := configuration.NewRedis(config)

	logger.InitLogger("")

	defer logger.Sync()

	//repository
	userRepository := repository.NewUserRepositoryImpl(database)
	cardRepository := repository.NewCardRepository(database)
	commentRepository := repository.NewCommentRepository(database)

	//service
	userService := service.NewUserServiceImpl(userRepository)
	cardService := service.NewCardServiceImpl(cardRepository)
	commentService := service.NewCommentServiceImpl(commentRepository)

	//controller
	userController := controller.NewUserController(&userService, config)
	cardController := controller.NewCardController(&cardService, config)
	commentController := controller.NewCommentController(&commentService, config)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())

	app.Use(recover.New())
	app.Use(cors.New())

	//swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//routing
	userController.Route(app)
	cardController.Route(app)
	commentController.Route(app)

	// Start the app
	//port := config.Get("SERVER.PORT")
	//if port == "" {
	//	log.Fatal("SERVER.PORT is not set in the configuration")
	//}

	logger.Info(context.Background(), "Server started on port 8080")
	if err := app.Listen(":8080"); err != nil {
		exception.PanicLogging(err)
	}
}
