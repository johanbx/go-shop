package page

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Page interface {
	BaseData() gin.H
	Template() string
}

func RenderPage(c *gin.Context, code int, p Page) {
	p.BaseData()["LIVE_RELOAD"] = os.Getenv("LIVE_RELOAD")
	c.HTML(code, p.Template(), p.BaseData())
}
