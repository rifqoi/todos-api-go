package main

import (
	"fmt"
	"os"

	_ "github.com/rifqoi/todos-api-go/docs"
	routes "github.com/rifqoi/todos-api-go/router"

	"github.com/gin-gonic/gin"
)

// @title Todo-API
// @version 1.0
// @description This is a API webservice to manage to-do list
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email hacktiv@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	router := gin.Default()
	routes.TodoRouters(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))

}
