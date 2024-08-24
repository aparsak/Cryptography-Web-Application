package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Serve static files (HTML, CSS)
	router.Static("/static", "./frontend")
	router.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	router.POST("/process", processCipher)

	router.Run(":8080")
}

func processCipher(c *gin.Context) {
	var requestData struct {
		Text   string `json:"text"`
		Key    string `json:"key"`
		Cipher string `json:"cipher"`
		Action string `json:"action"`
	}

	// Extract URL parameters
	cipher := c.Query("cipher")
	action := c.Query("action")

	// Bind JSON data to the requestData struct
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON or missing parameters"})
		return
	}

	// Override JSON fields with URL parameters if they are not present in the JSON body
	if requestData.Cipher == "" {
		requestData.Cipher = cipher
	}
	if requestData.Action == "" {
		requestData.Action = action
	}

	var result string
	var err error
	switch requestData.Cipher {
	case "caesar":
		result, err = caesarCipher(requestData.Text, requestData.Key, requestData.Action)
	case "affine":
		result, err = affineCipher(requestData.Text, requestData.Key, requestData.Action)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported cipher"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
