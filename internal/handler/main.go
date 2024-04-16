package handler

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.HTMLRender = renderer()

	r.POST("/catalog/item", PostCreateCatalogItem)
	r.GET("/catalog", GetAllCatalogItems)
	r.GET("/", PageIndex)
}

func renderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r = PageRenderer(r)
	return r
}
