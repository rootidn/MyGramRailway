package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)

		userRouter.PUT("", middlewares.Authentification(), controllers.UserUpdate)

		userRouter.DELETE("", middlewares.Authentification(), controllers.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentification())

		photoRouter.POST("", controllers.PhotoCreate)

		photoRouter.GET("", controllers.PhotoGetAll)

		photoRouter.GET("/:photoId", controllers.PhotoGetById)

		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.PhotoUpdate)

		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.PhotoDelete)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentification())

		commentRouter.POST("", controllers.CommentCreate)

		commentRouter.GET("", controllers.CommentGetAll)

		commentRouter.GET("/:commentId", controllers.CommentGetById)

		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.CommentUpdate)

		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.CommentDelete)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentification())

		socialMediaRouter.POST("", controllers.SocialMediaCreate)

		socialMediaRouter.GET("", controllers.SocialMediaGetAll)

		socialMediaRouter.GET("/:socialmediaId", middlewares.SocialMediaAuthorization(), controllers.SocialMediaGetById)

		socialMediaRouter.PUT("/:socialmediaId", middlewares.SocialMediaAuthorization(), controllers.SocialMediaUpdate)

		socialMediaRouter.DELETE("/:socialmediaId", middlewares.SocialMediaAuthorization(), controllers.SocialMediaDelete)
	}
	return r
}
