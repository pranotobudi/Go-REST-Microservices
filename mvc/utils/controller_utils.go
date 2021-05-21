package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}
