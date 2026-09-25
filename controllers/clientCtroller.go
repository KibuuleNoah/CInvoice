package controllers

import (
	"CInvoice/forms"
	"CInvoice/services"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	svc *services.ClientService
}

func NewClientController() *ClientController {
	return &ClientController{
		svc: services.NewClientService(),
	}
}

// GetClients godoc
// @Summary      Get all clients
// @Tags         clients
// @Produce      json
// @Param        limit   query     int  false  "Max results (default 10, max 50)"
// @Param        offset  query     int  false  "Pagination offset (default 0)"
// @Success      200     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]string
// @Router       /clients [get]
func (ctrl *ClientController) ListClients(c *gin.Context) {
	clients, err := ctrl.svc.ListClients(c.DefaultQuery("limit", "10"), c.DefaultQuery("offset", "0"))
	if err != nil {
		handleClientSeviceError(c, err)
		return
	}

	c.JSON(http.StatusOK, clients)
}

// GetClient godoc
// @Summary      Get a single client
// @Tags         clients
// @Produce      json
// @Param        id   path      string  true  "Client ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]string
// @Router       /clients/{id} [get]
func (ctrl *ClientController) GetClient(c *gin.Context) {
	id := c.Param("id")

	client, err := ctrl.svc.GetClientByID(id)
	if err != nil {
		handleClientSeviceError(c, err)
		return
	}

	c.JSON(http.StatusOK, client)
}

// CreateClient godoc
// @Summary      Create a client
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        body  body      forms.CreateClientForm  true  "Client details"
// @Success      201   {object}  map[string]interface{}
// @Failure      406   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /clients [post]
func (ctrl *ClientController) CreateClient(c *gin.Context) {
	var form forms.CreateClientForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		log.Println("CreateClient.Form Validation Error: ", validationErr)
		message := forms.Translate(validationErr, forms.CreateClientFormMessages)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": message})
		return
	}

	client, err := ctrl.svc.CreateClient(form)
	if err != nil {
		handleClientSeviceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, client)
}

// UpdateClient godoc
// @Summary      Update a client
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        body  body      forms.UpdateClientForm  true  "Client details"
// @Success      200   {object}  map[string]interface{}
// @Failure      406   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /clients [put]
func (ctrl *ClientController) UpdateClient(c *gin.Context) {
	var form forms.UpdateClientForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		log.Println("UpdateClient.Form Validation Error: ", validationErr)
		message := forms.Translate(validationErr, forms.UpdateClientFormMessages)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": message})
		return
	}

	// TODO: validate form
	client, err := ctrl.svc.UpdateClient(form)
	if err != nil {
		handleClientSeviceError(c, err)
		return
	}

	c.JSON(http.StatusOK, client)
}

// DeleteClient godoc
// @Summary      Delete a client
// @Tags         clients
// @Produce      json
// @Param        id   path      string  true  "Client ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /clients/{id} [delete]
func (ctrl *ClientController) DeleteClient(c *gin.Context) {
	id := c.Param("id")

	// TODO: validate form
	if err := ctrl.svc.DeleteClient(id); err != nil {
		handleClientSeviceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "client deleted"})
}

func handleClientSeviceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrClientNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	case errors.Is(err, services.ErrClientNameExists):
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
	case errors.Is(err, services.ErrClientCreate):
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	case errors.Is(err, services.ErrClientUpdate):
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	case errors.Is(err, services.ErrClientList):
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	case errors.Is(err, services.ErrClientDelete):
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong, please try again later"})
	}
}
