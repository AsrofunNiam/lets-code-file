package service

import (
	"io"
	"os"
	"strings"

	"github.com/AsrofunNiam/lets-code-file/exception"
	"github.com/AsrofunNiam/lets-code-file/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewProductService(
	db *gorm.DB,
	validate *validator.Validate,
) ProductService {
	return &ProductServiceImpl{
		DB:       db,
		Validate: validate,
	}
}

func (service *ProductServiceImpl) FindImage(fileName string, c *gin.Context) []byte {
	// Validate file name
	if !strings.HasSuffix(fileName, "jpg") {
		err := &exception.ErrorSendToResponse{Err: "Invalid file name. Only jpg files are allowed."}
		helper.PanicIfError(err)

	}

	// Open file
	fileProofPhoto := helper.PathToProduct + fileName
	fileOpen, err := os.Open(fileProofPhoto)
	helper.PanicIfError(err)

	fileBytes, err := io.ReadAll(fileOpen)
	helper.PanicIfError(err)

	return fileBytes
}
