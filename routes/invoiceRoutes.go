package routes

import (
	"CInvoice/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterInvoiceRoutes(r *gin.RouterGroup) {
	invoiceCtrl := controllers.NewInvoiceController()

	invoiceRoutes := r.Group("/invoices")
	{
		invoiceRoutes.GET("", invoiceCtrl.ListInvoices)
		invoiceRoutes.GET("/:id", invoiceCtrl.GetInvoice)
		invoiceRoutes.PATCH("/:id", invoiceCtrl.MarkInvoiceAsPaid)
	}

}
