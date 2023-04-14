package main

import (
	"mygram/controllers"
	"mygram/database"
	docs "mygram/docs"
	"mygram/routers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram
// @version 1.0
// @description MyGram is a free photo sharing app written in Go. People can share, view, and comment photos by everyone. Anyone can create an account by registering an email address and creating a username.
// @termOfService http://swagger.io/terms/

// @contact.name ilhamdy
// @contact.email ilhamdy.18@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description					        Description for what is this security definition being used

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	db := database.InitDB()

	docs.SwaggerInfo.BasePath = "/"

	userRepo := controllers.UserRepo{DB: db}
	photoRepo := controllers.PhotoRepo{DB: db}
	commentRepo := controllers.CommentRepo{DB: db}
	socialMediaRepo := controllers.SocialMediaRepo{DB: db}

	r := routers.StartApp(userRepo, photoRepo, commentRepo, socialMediaRepo)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
