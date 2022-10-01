package main

import (
	"github.com/KPWithCode/hurd/config"
	"github.com/KPWithCode/hurd/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	// use gin.Default() for router with default middleware
	router := gin.New()
	router.SetTrustedProxies(nil)
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}