package service

import (
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	FindImage(fileName string, c *gin.Context) []byte
}
