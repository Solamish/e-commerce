package router

import (

	"e-commerce/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LOAD(router *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	//middlerwares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(mw...)

	// 404 Handler.
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	//to check health state
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200})
	})

	//group
	digitalKeyGroup := router.Group("/api/")
	{

		digitalKeyGroup.POST("order", controller.PostForm)
		digitalKeyGroup.POST("addShop", controller.CreateShop)

	}
	return router
}
