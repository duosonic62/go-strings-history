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
	r.GET("/user", func(context *gin.Context) { webApp.UserQueryController.Show(context) })
	r.POST("/user", func(context *gin.Context) { webApp.UserCommandController.Create(context) })
	r.PUT("/user", func(context *gin.Context) { webApp.UserCommandController.Edit(context) })
	r.DELETE("/user", func(context *gin.Context) { webApp.UserCommandController.Delete(context) })

	r.Run(":8080")
}
