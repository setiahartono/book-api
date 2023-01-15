package server

import (
	"book-api/api/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	routerInstance := router.InitRouter(true)
	return routerInstance
}

func RunServer() {
	fmt.Println(`	
		██████   ██████   ██████  ██   ██                █████  ██████  ██ 
		██   ██ ██    ██ ██    ██ ██  ██                ██   ██ ██   ██ ██ 
		██████  ██    ██ ██    ██ █████       █████     ███████ ██████  ██ 
		██   ██ ██    ██ ██    ██ ██  ██                ██   ██ ██      ██ 
		██████   ██████   ██████  ██   ██               ██   ██ ██      ██				
	`)
	router := initRouter()
	router.Run("localhost:8080")
}
