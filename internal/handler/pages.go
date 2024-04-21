package handler

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func RootPageRenderer(r multitemplate.Renderer) {
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
}

func PageIndex(c *gin.Context) {
	TemplateResponse(c, http.StatusOK, "index", gin.H{
		"title": "Main website!",
	})
}
