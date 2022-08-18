package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/KPWithCode/hurd/controller"
) 

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.POST("/:id", controller.UpdateUser)
}
