package main

import (
	"encoding/gob"
	"johanbx/go-web-server-2/internal/catalog"
	"johanbx/go-web-server-2/internal/db"
	handler "johanbx/go-web-server-2/internal/handler"
	"net/http"
	"os"

	"github.com/gorilla/sessions"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(&catalog.CatalogItem{})
}

func main() {
	db.InitDB(os.Getenv("SQLITE_URI"))

	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.Use(SessionMiddleware())

	r.Static("/assets", "./assets")

	handler.PageRoutes(r)

	// Trigger reload on start-up
	if os.Getenv("GIN_MODE") != "release" && os.Getenv("LIVE_RELOAD") == "true" {
		http.Get("http://livereload:5555/trigger-reload")
	}

	r.Run("0.0.0.0:8080")
}

var store = sessions.NewCookieStore([]byte("replace-this-with-a-secure-random-key"))

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session-name")
		c.Set("session", session)
		c.Next()
	}
}
