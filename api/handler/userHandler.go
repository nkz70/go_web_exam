package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"webxam/controllers"
)

type UserHandler interface {
	FindUser(c *gin.Context)
}

type userHandler struct {
	UserController controllers.UserController
}

func NewUserHandler(uc controllers.UserController) UserHandler {
	return &userHandler{
		UserController: uc,
	}
}

func (uh *userHandler) FindUser(c *gin.Context) {
	var ur controllers.UserRequest
	if err := c.ShouldBindQuery(&ur); err != nil {
		log.Println(err)
	}

	users, err := uh.UserController.FindUser(&ur)
	if err != nil {

	}

	c.JSON(http.StatusOK, users)
}
