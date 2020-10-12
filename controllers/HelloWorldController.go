package controllers

import (
	"github.com/gin-gonic/gin"
)

type HelloWorldController struct{}

func (h *HelloWorldController) SampleHelloApi(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello world, climate change is real"})
}
