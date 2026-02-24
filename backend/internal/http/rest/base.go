package rest

import (
	"github.com/cg917658910/gotools/apiresp"
	"github.com/cg917658910/gotools/log"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	log.ZInfo(c, "Hello endpoint called")
	apiresp.GinSuccess(c, "Hello, World!")
}
