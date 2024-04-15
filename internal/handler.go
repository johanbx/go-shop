package handler

import (
	"net/http"
	"os"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.HTMLRender = renderer()

	r.GET("/ping", pong)
	r.GET("/", index)
}

func renderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	return r
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title":       "Main website!",
		"LIVE_RELOAD": os.Getenv("LIVE_RELOAD"),
	})
}
