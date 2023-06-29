package request

import "github.com/gin-gonic/gin"

var Request *gin.Context = nil

func SetRequest(c *gin.Context) {
	Request = c
}

func GetRequest() *gin.Context {
	return Request
}
