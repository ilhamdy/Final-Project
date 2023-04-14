package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp(u controllers.UserRepo, p controllers.PhotoRepo, c controllers.CommentRepo, s controllers.SocialMediaRepo) *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", u.UserRegister)
		userRouter.POST("/login", u.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", p.GetPhoto)
		photoRouter.GET("/:photoId", middlewares.PhotoAuthorization(), p.GetPhotoById)
		photoRouter.POST("/", p.CreatePhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), p.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), p.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", c.GetComment)
		commentRouter.GET("/:commentId", middlewares.CommentAuthorization(), c.GetCommentById)
		commentRouter.POST("/", c.CreateComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), c.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), c.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedia")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.GET("/", s.GetSocialMedia)
		socialMediaRouter.GET("/:socialMediaId", middlewares.SocialMediaAuthorization(), s.GetSocialMediaById)
		socialMediaRouter.POST("/", s.CreateSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), s.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), s.DeleteSocialMedia)
	}

	return r
}
