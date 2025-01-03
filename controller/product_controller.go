package controller

import (
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindImage(context *gin.Context)
}
