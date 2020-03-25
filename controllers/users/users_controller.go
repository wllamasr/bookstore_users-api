package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/wllamasr/bookstore_users-api/domain/users"
	"github.com/wllamasr/bookstore_users-api/services"
	"github.com/wllamasr/bookstore_users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if error := c.ShouldBindJSON(&user); error != nil {
		restError := errors.BadRequestError("Invalid body JSON")
		c.JSON(restError.Status, restError)
		return
	}

	result, saveError := services.CreateUser(user)

	if saveError != nil {
		//TODO: Implement saveError
		c.JSON(saveError.Status, saveError)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
