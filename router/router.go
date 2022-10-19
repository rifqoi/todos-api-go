package routes

import (
	"github.com/rifqoi/todos-api-go/controllers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func TodoRouters(router *gin.Engine) {
	router.GET("/todos", controllers.GetAllTodos)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
