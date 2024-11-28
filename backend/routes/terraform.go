package routes

import (
	"terraform-manager/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/terraform/init", controllers.InitTerraform)
		api.POST("/terraform/apply", controllers.ApplyTerraform)
		api.POST("/terraform/destroy", controllers.DestroyTerraform)
	}
}
