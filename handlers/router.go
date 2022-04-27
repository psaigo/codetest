package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(svr Service) *gin.Engine {

	// Create service instance. Main instance shared by all http handlers.
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Group("/cryptoserver")
	{
		v1Group := router.Group("/cryptoserver/v1")
		{
			v1Group.GET("/currency/:symbol", svr.GetCurrency)
			v1Group.GET("/currency/all", svr.GetAllCurrency)
		}
	}
	return router
}
