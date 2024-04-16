package handler

import (
	"net/http"
	"os"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func PageRenderer(renderer multitemplate.Renderer) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	return r
}

func PageIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title":       "Main website!",
		"LIVE_RELOAD": os.Getenv("LIVE_RELOAD"),
	})
}
