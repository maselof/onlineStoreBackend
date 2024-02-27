package store

import (
	"github.com/gin-gonic/gin"
	"log"
	"onlineStoreBackend/api/routes"
	"onlineStoreBackend/constants"
)

func Run() {
	router := gin.Default()

	var restRoutes routes.RestRoutes
	restRoutes.Setup(router.Group(constants.BaseRoutePrefix))

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
