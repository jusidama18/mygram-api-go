package main

import (
	_ "github.com/jusidama18/mygram-api-go/docs"
	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/controllers"
	"github.com/jusidama18/mygram-api-go/api/middlewares"
	"github.com/jusidama18/mygram-api-go/api/routes"
	"github.com/jusidama18/mygram-api-go/config"
	"github.com/jusidama18/mygram-api-go/repository/gorm"
	"github.com/jusidama18/mygram-api-go/services"
)


// @title Mygram-API
// @version 1.0
// @description This is a API webservice to manage mygram API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email hacktiv@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic(err)
	}

	userRepo := gorm.NewUserRepository(db)
	userServices := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userServices)

	photoRepo := gorm.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)

	commentRepo := gorm.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	commentController := controllers.NewCommentController(commentService)

	socialMediaRepo := gorm.NewSocialMediaRepository(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	middleware := middlewares.NewMiddleware(userServices)

	router := gin.Default()
	app := routes.NewRouter(router, userController, photoController, commentController, socialMediaController, middleware)
	app.Run()
}
