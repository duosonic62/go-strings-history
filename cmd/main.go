package main

import (
	"github.com/duosonic/go-strings-history/internal/adaptor/infrastructure/router"
	"github.com/gin-gonic/gin"
)

func main() {
	func main() {
		r := router.Router
		var webApp = injector.Initialize()

		// ユーザ登録API
		r.POST("/user", func(context *gin.Context) { webApp.UserController.CreateUser(context) })

		r.Run(":8080")
	}
}
