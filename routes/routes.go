package routes

import (
	"golang-frontend/handler"
	"golang-frontend/middleware"

	"github.com/gin-gonic/gin"
)

func RoutesHtml(router *gin.RouterGroup) {
	router.Use(middleware.JwtAuthHtmlMiddleware())
	router.GET("/welcome", handler.Welcome)
	router.GET("/schedule", handler.Schedule)
	router.GET("/viewSchedule", handler.ViewSchedule)
	router.GET("/editSchedule", handler.EditSchedule)
	router.GET("/advanceSetting", handler.AdvanceSetting)
}
func RoutesJson(router *gin.RouterGroup) {
	router.Use(middleware.JwtAuthJsonMiddleware())
	router.POST("/logout", handler.Logout)
	router.POST("/schemaList", handler.SchemaList)
}
