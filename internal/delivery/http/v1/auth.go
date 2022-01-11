package httpv1

import (
	"golang-template-module/internal/models"
	"golang-template-module/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth godoc
// @Summary      Auth user
// @Description  auth User
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Auth   body   models.UserAuthModel  true "Auntification"
// @Success      200  {object}  models.UserTokenModel
// @Router       /auth [post]
func RouterAuth(c *gin.Context) error {
	user := new(models.UserAuthModel)
	if err := c.ShouldBindJSON(user); err != nil {
		return err
	}
	tokenModel, err := service.Authorization(user)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, tokenModel)
	return nil
}

// CreateUser godoc
// @Summary      Create user
// @Description  create user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Create   body   models.UserAuthModel  true "Creating"
// @Success      200  {object}  models.UserTokenModel
// @Router       /create [post]
func RouterRegistration(c *gin.Context) error {
	user := new(models.UserAuthModel)
	if err := c.ShouldBindJSON(user); err != nil {
		return err
	}
	tokenModel, err := service.CreateUser(user)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, tokenModel)
	return nil
}
