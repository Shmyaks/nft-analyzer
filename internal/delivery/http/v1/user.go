package httpv1

import (
	"golang-template-module/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserGet godoc
// @Summary      Get user
// @Description  Get user by id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "id"
// @Success      200  {object}  models.UserBaseModel
// @Router       /user/{id} [get]
func RouterUserGetId(c *gin.Context) error {
	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	user, err := service.GetUserService(u64)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, *user)
	return nil
}

// UserGetList godoc
// @Summary      Get user list
// @Description  Get user list
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UserListBaseModel
// @Router       /user/list [get]
func RouterUserGetList(c *gin.Context) error {
	users, err := service.GetUserListService()
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, users)
	return nil
}
