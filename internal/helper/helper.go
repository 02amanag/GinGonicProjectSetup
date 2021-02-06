//this is very important sector where all the reusable aur small distinguish function can be wrriten here
package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayError(c *gin.Context, message string, status int) {
	print(message + " <<<- this error @ this endpoint ->>> ") // to print all erroe at console
	if message == "Section not Awailable" {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Title": "Unacceptable", "Message": message, "Status": 406})
		return
	}
	if status == 401 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Title": "Please login", "Message": message, "Status": status})
	} else {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": message, "Status": status})
	}
}
