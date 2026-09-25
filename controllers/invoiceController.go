package controllers

import (
	"CInvoice/services"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvoiceController struct {
	svc *services.InvoiceService
}

func NewInvoiceController() *InvoiceController {
	return &InvoiceController{svc: services.NewInvoiceService()}
}

// GetInvoices godoc
// @Summary      Get all invoices
// @Tags         invoices
// @Produce      json
// @Param        limit   query     int  false  "Max results (default 10, max 50)"
// @Param        offset  query     int  false  "Pagination offset (default 0)"
// @Success      200     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]string
// @Router       /invoices [get]
func (ctrl *InvoiceController) ListInvoices(c *gin.Context) {
	invoices, err := ctrl.svc.ListInvoices(c.DefaultQuery("limit", "10"), c.DefaultQuery("offset", "0"))
	if err != nil {
		handleInvoiceSeviceError(c, err)
		return
	}
	c.JSON(http.StatusOK, invoices)
}

// GetInvoice godoc
// @Summary      Get a single invoice
// @Tags         invoices
// @Produce      json
// @Param        id   path      string  true  "Invoice ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]string
// @Router       /invoices/{id} [get]
func (ctrl *InvoiceController) GetInvoice(c *gin.Context) {
	id := c.Param("id")

	invoice, err := ctrl.svc.GetInvoiceByID(id)
	if err != nil {
		handleInvoiceSeviceError(c, err)
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func (ctrl *InvoiceController) MarkInvoiceAsPaid(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.svc.MarkInvoiceAsPaid(id); err != nil {
		handleInvoiceSeviceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invoice marked as paid"})
}

func (ctrl *InvoiceController) CreateClientInvoice(c *gin.Context) {
	clientID := c.Param("id")

	if err := ctrl.svc.CreateClientInvoice(clientID); err != nil {
		log.Println("InvoiceController.CreateClientInvoice ERROR: ", err)
		handleInvoiceSeviceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invoice created"})
}

func handleInvoiceSeviceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrInvoiceNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	case errors.Is(err, services.ErrInvoiceCreate):
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	case errors.Is(err, services.ErrInvoiceUpdate):
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	case errors.Is(err, services.ErrInvoiceList):
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	case errors.Is(err, services.ErrInvoiceDelete):
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	case errors.Is(err, services.ERRInvoiceExists):
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong, please try again later"})
	}
}
