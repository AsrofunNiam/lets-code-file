package route

import (
	"github.com/AsrofunNiam/lets-code-file/controller"
	"github.com/AsrofunNiam/lets-code-file/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ProductRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {
	Products := service.NewProductService(
		db,
		validate,
	)
	productController := controller.NewProductController(Products)
	router.GET("/products/photo/:file_name", productController.FindImage)

}
