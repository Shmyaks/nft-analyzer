package handler

import (
	"errors"
	"golang-template-module/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(myhandler func(c *gin.Context) error) (ginhandler func(c *gin.Context)) {
	return func(c *gin.Context) {
		var appErr *models.AppError
		if err := myhandler(c); err != nil {
			if errors.As(err, &appErr) {
				c.JSON(appErr.StatusCode, appErr)
			} else {
				c.JSON(http.StatusBadRequest, err.Error())
			}
		}
	}
}
