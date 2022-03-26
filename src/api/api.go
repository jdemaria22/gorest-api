package routes

import (
	"os"

	"github.com/gin-gonic/gin"

	"demoproject/src/docs"

	service "demoproject/src/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitialiseRoutes() {
	docs.SwaggerInfo.Description = os.Getenv("DESCRIPTION")
	docs.SwaggerInfo.Title = os.Getenv("ARTIFACT")
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.Version = os.Getenv("VERSION")
	r := gin.Default()

	user := r.Group("/user")
	{
		user.POST("", service.PostUser)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + os.Getenv("PORT"))
}
