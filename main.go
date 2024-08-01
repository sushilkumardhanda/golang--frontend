package main

import (
	"golang-frontend/handler"
	"golang-frontend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./web/assets")
	router.LoadHTMLGlob("web/templates/*")

	router.GET("/login", handler.Login)
	router.POST("/login", handler.LoginPost)
	router.POST("/loginConfirm", handler.LoginConfirm)
	html := router.Group("")
	routes.RoutesHtml(html)
	json := router.Group("")
	routes.RoutesJson(json)
	router.Run(":8000")

}
