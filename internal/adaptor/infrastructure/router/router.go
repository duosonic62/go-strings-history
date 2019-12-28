package router

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init() {
	router := gin.Default()

	Router = router
}
