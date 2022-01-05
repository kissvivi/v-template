package router

import (
	"v-template/internal/controller"
	"github.com/gin-gonic/gin"
)

func routers(r *gin.Engine, controller *controller.Controller) {
	g := r.Group("/test")
	{
		g.GET("/ok", controller.TestOk)
		g.GET("/demo/ok", controller.DemoOk)
		g.POST("/create", controller.CreateTest)
	}
}
