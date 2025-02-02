package main

import (
	"github.com/BiswajitThakur/toy-ses/api/email"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	router := gin.Default()

	router.Use(EmailHandler())

	v1 := router.Group("/api/email")
	{
		v1.GET("/count", email.CountEmail)
		v1.POST("/send/:email", email.SendEmail)
	}

	return router
}
