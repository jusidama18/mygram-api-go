package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jusidama18/mygram-api-go/api/controllers"
	"github.com/jusidama18/mygram-api-go/api/middlewares"
)

type Router struct {
	router      *gin.Engine
	user        *controllers.UserController
	photo       *controllers.PhotoController
	comment     *controllers.CommentController
	socialMedia *controllers.SocialMediaController
	middleware  *middlewares.Middleware
}

func NewRouter(router *gin.Engine, user *controllers.UserController, photo *controllers.PhotoController, comment *controllers.CommentController, socialMedia *controllers.SocialMediaController, middleware *middlewares.Middleware) *Router {
	return &Router{
		router:      router,
		user:        user,
		photo:       photo,
		comment:     comment,
		socialMedia: socialMedia,
		middleware:  middleware,
	}
}

func (r *Router) Run() {
	userRoutes := r.router.Group("/users")
	userRoutes.POST("/register", r.user.RegisterUser)
	userRoutes.POST("/login", r.user.Login)
	userRoutes.PUT("/", r.middleware.Authorization, r.user.UpdateUser)
	userRoutes.DELETE("/", r.middleware.Authorization, r.user.DeleteUser)

	photoRoutes := r.router.Group("/photos").Use(r.middleware.Authorization)
	photoRoutes.POST("/", r.photo.CreatePhoto)
	photoRoutes.GET("/", r.photo.GetAllPhotos)
	photoRoutes.PUT("/:id", r.photo.UpdatePhoto)
	photoRoutes.DELETE("/:id", r.photo.DeletePhoto)

	commentRoutes := r.router.Group("/comments").Use(r.middleware.Authorization)
	commentRoutes.POST("/", r.comment.CreateComment)
	commentRoutes.GET("/", r.comment.GetAllComment)
	commentRoutes.PUT("/:id", r.comment.UpdateComment)
	commentRoutes.DELETE("/:id", r.comment.DeleteComment)

	socialMediaRoutes := r.router.Group("/socialmedias").Use(r.middleware.Authorization)
	socialMediaRoutes.POST("/", r.socialMedia.CreateSocialMedia)
	socialMediaRoutes.GET("/", r.socialMedia.GetAllSocialMedia)
	socialMediaRoutes.PUT("/:id", r.socialMedia.UpdateSocialMedia)
	socialMediaRoutes.DELETE("/:id", r.socialMedia.DeleteSocialMedia)

	// Check middleware
	r.router.GET("/check", r.middleware.Authorization, r.user.Check)

	r.router.Run(":8080")
}
