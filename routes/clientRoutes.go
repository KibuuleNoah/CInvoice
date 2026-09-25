package routes

import (
	"CInvoice/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterClientRoutes(r *gin.RouterGroup) {
	clientCtrl := controllers.NewClientController()

	clientRoutes := r.Group("/clients")
	{
		clientRoutes.GET("", clientCtrl.ListClients)
		clientRoutes.GET("/:id", clientCtrl.GetClient)
		clientRoutes.POST("", clientCtrl.CreateClient)
		clientRoutes.PUT("/:id", clientCtrl.UpdateClient)
		clientRoutes.DELETE("/:id", clientCtrl.DeleteClient)
	}
}
