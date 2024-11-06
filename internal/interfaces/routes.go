package interfaces

import (
    "github.com/gin-gonic/gin"
    "kirito/internal/usecase"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"msg": usecase.GetHelloWorldMessage()})
    })
	
	router.GET("/getme", func(c *gin.Context) {
        c.JSON(200, usecase.GetDummyUser())
    })

	
    return router
}