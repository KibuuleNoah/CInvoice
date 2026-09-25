package controllers

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebController struct{}

func NewWebController(router *gin.Engine, embededTemplatesFS embed.FS) *WebController {
	router.SetFuncMap(template.FuncMap{
		"inc":      func(i int) int { return i + 1 },
		"safeHTML": func(s template.HTML) template.HTML { return s },
		"default": func(defaultValue, val interface{}) interface{} {
			if val == nil || val == "" {
				return defaultValue
			}
			return val
		},
	})

	// router.LoadHTMLGlob("web/templates/*.html")
	router.LoadHTMLFS(http.FS(embededTemplatesFS), "web/templates/*.html")

	// ── Page routes ──
	router.GET("/clients", Clients)
	router.GET("/invoices", Invoices)

	return &WebController{}
}

func Clients(c *gin.Context) {
	c.HTML(http.StatusOK, "clients", nil)
}

func Invoices(c *gin.Context) {
	c.HTML(http.StatusOK, "invoices", nil)
}
