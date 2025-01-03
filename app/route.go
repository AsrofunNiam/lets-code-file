package app

import (
	"fmt"
	"runtime/debug"

	"github.com/AsrofunNiam/lets-code-file/exception"
	route "github.com/AsrofunNiam/lets-code-file/route"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// ErrorHandler
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
				exception.ErrorHandler(c, err)
			}
		}()
		c.Next()
	}
}

func NewRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {

	// Set gin to release mode
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	//  exception middleware
	router.Use(ErrorHandler())
	router.UseRawPath = true

	// route path
	route.ProductRoute(router, db, validate)

	return router
}
