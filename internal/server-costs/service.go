package servercosts

import (
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

const (
	apiTokenHeader  = "API_TOKEN"
	errorMessageKey = "error_message"
)

func ServerCosts(c *gin.Context) {

	apiToken := c.GetHeader(apiTokenHeader)

	response, err := serverCosts(apiToken)
	if err != nil {
		c.JSON(response.errorCode, gin.H{
			errorMessageKey: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, structs.Map(&response))
}

