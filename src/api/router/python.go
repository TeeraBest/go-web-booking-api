package router

import (
	"go-web-api/api/handler"

	"github.com/gin-gonic/gin"
)

func Python(g *gin.RouterGroup) {
	handler := handler.NewPythonHandler()
	g.GET("/hello", handler.Hello)
	g.POST("/predict", handler.Predict)
}
