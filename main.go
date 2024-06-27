package main

import (
	"net/http"

	"golangchain/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/chat/", api.ChatHandler)

	r.Run(":7860")
}
