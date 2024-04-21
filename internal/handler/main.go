package handler

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func PageRoutes(r *gin.Engine) {
	r.HTMLRender = renderer()

	r.POST("/catalog/item", CatalogItemHandler)
	r.GET("/catalog/item", CatalogItemHandler)

	r.GET("/", PageIndex)
}

func renderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	RootPageRenderer(r)
	CatalogPageRenderer(r)
	return r
}
