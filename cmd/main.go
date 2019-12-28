package main

import (
	"github.com/duosonic62/go-strings-history/cmd/injector"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/db/migration"
	"github.com/duosonic62/go-strings-history/internal/adaptor/infrastructure/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := router.Router
	webApp := injector.Initialize()

	// db settings
	migration.Migrate()
	db.InitDB()

	// ユーザ登録API
	r.POST("/user", func(context *gin.Context) { webApp.UserController.CreateUser(context) })

	r.Run(":8080")
}
