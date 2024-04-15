package main

import (
	handler "johanbx/go-web-server-2/internal"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.Static("/assets", "./assets")

	handler.Register(r)

	// Trigger reload on start-up
	if os.Getenv("GIN_MODE") != "release" && os.Getenv("LIVE_RELOAD") == "true" {
		http.Get("http://livereload:5555/trigger-reload")
	}

	r.Run("0.0.0.0:8080")
}
