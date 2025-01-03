package controller

import (
	"github.com/AsrofunNiam/lets-code-file/service"
	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) FindImage(c *gin.Context) {
	fileName := c.Param("file_name")
	fileResponse := controller.ProductService.FindImage(fileName, c)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Write(fileResponse)
}
