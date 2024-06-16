package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ChatHandler(c *gin.Context) {
	var inputData ChatInput
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	apiKey := os.Getenv("GROQ_API_KEY")
	response, err := InvokeChain(apiKey, inputData.UserQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response := "Hello, World!"
	c.JSON(http.StatusOK, response)
}
